package util

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// Rand

func RandUint(min uint, max uint) uint {
	if min > max {
		min, max = max, min
	}
	return (uint(rand.Int63()) % (max - min + 1)) + min
}

// Diffie-Hellman Key-exchange
var (
	rng         = rand.New(rand.NewSource(time.Now().UnixNano()))
	DH1BASE     = big.NewInt(3)
	DH1PRIME, _ = big.NewInt(0).SetString("0x7FFFFFC3", 0)
	MAXINT64    = big.NewInt(math.MaxInt64)
)

func Itob(b *big.Int) []byte {
	return []byte(fmt.Sprintf("%v", b))
}

func DHExchange() (*big.Int, *big.Int) {
	secret := big.NewInt(0).Rand(rng, MAXINT64)
	modpower := big.NewInt(0).Exp(DH1BASE, secret, DH1PRIME)
	return secret, modpower
}

func DHKey(SECRET, MODPOWER *big.Int) *big.Int {
	key := big.NewInt(0).Exp(MODPOWER, SECRET, DH1PRIME)
	return key
}

// Padding
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
