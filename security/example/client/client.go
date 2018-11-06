package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	// hand shake
	rsa := util.NewRsaCipher("pem/server_public.pem", "pem/client_private.pem")
	key0 := []byte(fmt.Sprintf("%v", util.RandUint(1000, 9999)))
	msg0 := msg.MsgHandShake{}
	msg0.Key = rsa.Encrypt(key0)
	bts, _ := json.Marshal(&msg0)
	conn.Write(bts)
	conn.Write([]byte("\n"))
	log.Printf("key0:%s,Key:%s", key0, msg0.Key)
	bts, _, _ = reader.ReadLine()
	msg00 := msg.MsgHandShake{}
	json.Unmarshal(bts, &msg00)
	key00 := rsa.Decrypt(msg00.Key)
	log.Printf("key0:%s,key00:%s", key0, key00)
	if string(key0) != string(key00) {
		log.Printf("hand shake fail")
		return
	}

	// key exchange
	x1, e1 := util.DHExchange()
	x2, e2 := util.DHExchange()
	msg1 := msg.MsgKeyExchange{}
	msg1.SendSeed = e1.Int64()
	msg1.ReciveSeed = e2.Int64()
	bts, _ = json.Marshal(&msg1)
	conn.Write(bts)
	conn.Write([]byte("\n"))
	log.Printf("send key exchange:%s", bts)

	bts, _, _ = reader.ReadLine()
	json.Unmarshal(bts, &msg1)
	log.Printf("recive key exchange %v", msg1)

	key1 := util.DHKey(x1, big.NewInt(msg1.SendSeed))
	key2 := util.DHKey(x2, big.NewInt(msg1.ReciveSeed))
	log.Printf("key1:%v,key2:%v", key1, key2)

	// data encode/decode
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
