package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/big"
	"net"

	"github.com/MaxwellBackend/Games/security"
	"github.com/MaxwellBackend/Games/security/example/msg"
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
	// hand shake
	rsa := security.NewRsaCipher("pem/client_public.pem", "pem/server_private.pem")
	line, _, _ := reader.ReadLine()
	msg0 := msg.MsgHandShake{}
	json.Unmarshal(line, &msg0)
	key0 := rsa.Decrypt(msg0.Key)
	msg0.Key = rsa.Encrypt(key0)
	log.Printf("key0:%s,Key:%s", key0, msg0.Key)
	line, _ = json.Marshal(&msg0)
	conn.Write(line)
	conn.Write([]byte("\n"))

	// key exchange
	line, _, _ = reader.ReadLine()
	msg1 := msg.MsgKeyExchange{}
	json.Unmarshal(line, &msg1)
	log.Printf("recive key exchange  %v", msg1)

	x1, e1 := security.DHExchange()
	key1 := security.DHKey(x1, big.NewInt(msg1.SendSeed))
	x2, e2 := security.DHExchange()
	key2 := security.DHKey(x2, big.NewInt(msg1.ReciveSeed))
	log.Printf("key1:%v,key2:%v", key1, key2)

	msg1.SendSeed = e1.Int64()
	msg1.ReciveSeed = e2.Int64()
	line, _ = json.Marshal(&msg1)
	conn.Write(line)
	conn.Write([]byte("\n"))
	log.Printf("response key exchange  %s", line)

	// data encode/decode
	decoder := security.NewAesCipher(security.Itob(key1))
	encoder := security.NewAesCipher(security.Itob(key2))

	line, _, _ = reader.ReadLine()
	line = decoder.Decrypt(line)
	log.Printf("decrypt data %s\n", line)

	conn.Write(encoder.Encrypt(line))
	conn.Write([]byte("\n"))
}
