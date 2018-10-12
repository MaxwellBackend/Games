package main

import (
	"crypto/cipher"
	"crypto/des"
)

var desblock cipher.Block

func init() {
	var key [8]byte
	for i := range key {
		key[i] = byte(RandUint(0, 255))
	}
	desblock, _ = des.NewCipher(key[:])
}

func DesEncrypt(buf, encrypted []byte) {
	desblock.Encrypt(buf, encrypted)
}

func DesDecrypt(encrypted, buf []byte) {
	desblock.Decrypt(encrypted, buf)
}
