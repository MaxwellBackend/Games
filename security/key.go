package main

import (
	"crypto/aes"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"crypto/rsa"
	"crypto/sha1"
)

var KEY []byte
var KEY8 [8]byte
var KEY32 [32]byte

func init() {
	/*x1, _ := DHExchange()
	key1 := DHKey(x1, big.NewInt(int64(123)))
	KEY = []byte(fmt.Sprintf("%v", key1))
	*/
	pubk, _ = PubKey()
	prik, _ = PriKey()
	var buf [128]byte
	KEY, _ = rsa.EncryptOAEP(sha1.New(), rand.Reader, pubk, []byte(buf[:]), []byte(""))

	copy(KEY8[:], KEY)
	copy(KEY32[:], KEY)

	/*
		fmt.Println("key length:", len(KEY), string(KEY), KEY)
		fmt.Println("key8 length:", len(KEY8), string(KEY8[:]), KEY8)
		fmt.Println("key32 length:", len(KEY32), string(KEY32[:]), KEY32)
	*/
	aesblock, _ = aes.NewCipher(KEY32[:])
	desblock, _ = des.NewCipher(KEY8[:])
	rc4cipher, _ = rc4.NewCipher(KEY32[:])
}

var pubk *rsa.PublicKey
var prik *rsa.PrivateKey
