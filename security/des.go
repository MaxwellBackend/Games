package main

import (
	"crypto/cipher"
)

var desblock cipher.Block

func DesEncrypt(src []byte) []byte {
	blockSize := desblock.BlockSize()
	src = PKCS5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(desblock, KEY32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst

}

func DesDecrypt(src []byte) []byte {
	blockSize := desblock.BlockSize()
	blockMode := cipher.NewCBCDecrypter(desblock, KEY32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = PKCS5UnPadding(dst)
	return dst

}
