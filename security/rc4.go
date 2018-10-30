package main

import (
	"crypto/rc4"
)

var erc4cipher *rc4.Cipher
var drc4cipher *rc4.Cipher

func Rc4Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	erc4cipher.XORKeyStream(dst, src)
	return dst
}

func Rc4Decrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	drc4cipher.XORKeyStream(dst, src)
	return dst
}
