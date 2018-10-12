package main

import (
	"crypto/cipher"
)

var desblock cipher.Block

func DesEncrypt(buf, encrypted []byte) {
	desblock.Encrypt(buf, encrypted)
}

func DesDecrypt(encrypted, buf []byte) {
	desblock.Decrypt(encrypted, buf)
}
