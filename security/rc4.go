package main

import (
	"crypto/rc4"
)

var rc4cipher *rc4.Cipher

func Rc4Encrypt(buf, encrypted []byte) {
	rc4cipher.XORKeyStream(buf, encrypted)
}

func Rc4Decrypt(encrypted, buf []byte) {
	rc4cipher.XORKeyStream(encrypted, buf)
}
