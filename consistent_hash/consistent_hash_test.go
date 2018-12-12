package consistent_hash

import (
	"testing"
	"fmt"
)

func TestConsistentHash(t *testing.T) {

	ch := NewConsistentHashExecutor()
	ch.AddNode("a")

	ch.AddNode("b")
	ch.AddNode("c")
	ch.AddNode("d")

	fmt.Println(ch.Lookup("golang"))
	fmt.Println(ch.Lookup("java"))
	fmt.Println(ch.Lookup("php"))
	fmt.Println(ch.Lookup("c"))
	fmt.Println(ch.Lookup("cpp"))
	fmt.Println(ch.Lookup("kotlin"))
	fmt.Println(ch.Lookup("c#"))
}
