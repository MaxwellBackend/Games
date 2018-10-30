package main

import (
	"crypto/cipher"
)

var aesblock cipher.Block

func AesEncrypt(src []byte) []byte {
	blockSize := aesblock.BlockSize()
	src = PKCS5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(aesblock, KEY32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	// aesblock.Encrypt(dst, src)
	return dst
}

func AesDecrypt(src []byte) []byte {
	blockSize := aesblock.BlockSize()
	blockMode := cipher.NewCBCDecrypter(aesblock, KEY32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	// aesblock.Decrypt(dst, src)
	dst = PKCS5UnPadding(dst)
	return dst
}
