package consistent_hash

import (
	"hash/crc32"
	"fmt"
	"errors"
	"sort"
)

const VIRTUAL_NODE = 64 // 每个物理节点包含的虚拟节点个数

// 一致性hash执行器
type ConsistentHashExecutor struct {
	physicalNodes map[string]bool // 物理节点
	virtualNodes  []*VirtualNode  // 虚拟节点
}

// 虚拟节点
type VirtualNode struct {
	name             string // 节点名称
	hashcode         uint32 // hash值
	physicalNodeName string // 物理节点名称
}

// 创建一致性hash执行器
func NewConsistentHashExecutor() *ConsistentHashExecutor {
	return &ConsistentHashExecutor{
		physicalNodes: make(map[string]bool),
	}
}

func (ch *ConsistentHashExecutor) Len() int {
	return len(ch.virtualNodes)
}

func (ch *ConsistentHashExecutor) Swap(i, j int) {
	ch.virtualNodes[i], ch.virtualNodes[j] = ch.virtualNodes[j], ch.virtualNodes[i]
}

func (ch *ConsistentHashExecutor) Less(i, j int) bool {
	return ch.virtualNodes[i].hashcode < ch.virtualNodes[j].hashcode
}

// 添加物理节点
func (ch *ConsistentHashExecutor) AddNode(name string) error {

	// 物理节点已添加
	if ch.physicalNodes[name] {
		return errors.New("node already add")
	}

	// 新增物理节点
	ch.physicalNodes[name] = true

	// 新增虚拟节点
	for i := 1; i <= VIRTUAL_NODE; i++ {

		virtualNodeKey := fmt.Sprintf("%v_%v", name, i)
		table := crc32.MakeTable(crc32.IEEE)
		hashcode := crc32.Checksum([]byte(virtualNodeKey), table)

		ch.virtualNodes = append(ch.virtualNodes, &VirtualNode{
			name:             virtualNodeKey,
			hashcode:         hashcode,
			physicalNodeName: name,
		})
	}

	sort.Sort(ch)

	return nil
}

// 删除物理节点
func (ch *ConsistentHashExecutor) DelNode(name string) {

	// 删除虚拟节点
	for k, v := range ch.virtualNodes {
		if v.physicalNodeName == name {
			ch.virtualNodes = append(ch.virtualNodes[:k], ch.virtualNodes[k+1:]...)
		}
	}

	// 删除物理节点
	delete(ch.physicalNodes, name)
}

// 查找key所在的节点
func (ch *ConsistentHashExecutor) Lookup(key string) string {

	table := crc32.MakeTable(crc32.IEEE)
	hashcode := crc32.Checksum([]byte(key), table)

	if len(ch.virtualNodes) == 0 {
		fmt.Println("lookup virtualNodes is empty")
		return ""
	}

	for _, vn := range ch.virtualNodes {
		if vn.hashcode >= hashcode {
			return vn.physicalNodeName
		}
	}

	return ch.virtualNodes[0].physicalNodeName
}
