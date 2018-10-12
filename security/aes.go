package main

import (
	"crypto/cipher"
)

var aesblock cipher.Block

func AesEncrypt(buf, encrypted []byte) {
	aesblock.Encrypt(buf, encrypted)
}

func AesDecrypt(encrypted, buf []byte) {
	aesblock.Decrypt(encrypted, buf)
}
