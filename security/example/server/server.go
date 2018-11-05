package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/big"
	"net"
	"security/example/msg"
	"security/util"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Printf("listen err %v\n", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err %v\n", err)
			break
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	line, _, _ := reader.ReadLine()

	msg1 := msg.MsgKeyExchange{}
	json.Unmarshal(line, &msg1)
	log.Printf("recive key exchange  %v", msg1)

	x1, e1 := util.DHExchange()
	key1 := util.DHKey(x1, big.NewInt(msg1.SendSeed))
	x2, e2 := util.DHExchange()
	key2 := util.DHKey(x2, big.NewInt(msg1.ReciveSeed))
	log.Printf("key1:%v,key2:%v", key1, key2)

	msg1.SendSeed = e1.Int64()
	msg1.ReciveSeed = e2.Int64()
	line, _ = json.Marshal(&msg1)
	conn.Write(line)
	conn.Write([]byte("\n"))
	log.Printf("response key exchange  %s", line)

	decoder := util.NewAesCipher(util.Itob(key1))
	encoder := util.NewAesCipher(util.Itob(key2))

	line, _, _ = reader.ReadLine()
	line = decoder.Decrypt(line)
	log.Printf("decrypt data %s\n", line)

	conn.Write(encoder.Encrypt(line))
	conn.Write([]byte("\n"))
}
