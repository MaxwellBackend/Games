package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/big"
	"net"
	"security/example/msg"
	"security/util"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Printf("dial err %v\n", err)
		return
	}
	reader := bufio.NewReader(conn)
	x1, e1 := util.DHExchange()
	x2, e2 := util.DHExchange()

	msg1 := msg.MsgKeyExchange{}
	msg1.SendSeed = e1.Int64()
	msg1.ReciveSeed = e2.Int64()
	bts, _ := json.Marshal(&msg1)
	conn.Write(bts)
	conn.Write([]byte("\n"))
	log.Printf("send key exchange:%s", bts)

	bts, _, _ = reader.ReadLine()
	json.Unmarshal(bts, &msg1)
	log.Printf("recive key exchange %v", msg1)

	key1 := util.DHKey(x1, big.NewInt(msg1.SendSeed))
	key2 := util.DHKey(x2, big.NewInt(msg1.ReciveSeed))

	log.Printf("key1:%v,key2:%v", key1, key2)

	encoder := util.NewAesCipher(util.Itob(key1))
	decoder := util.NewAesCipher(util.Itob(key2))

	msg2 := msg.MsgHello{}
	msg2.Data = "hello world!"
	bts, _ = json.Marshal(&msg2)
	conn.Write(encoder.Encrypt(bts))
	conn.Write([]byte("\n"))

	bts, _, _ = reader.ReadLine()
	bts = decoder.Decrypt(bts)
	log.Printf("decrypt data %s\n", bts)

	time.Sleep(3 * 1e9)
}
