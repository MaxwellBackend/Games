package main

import (
	"fmt"
	"testing"
)

var bts []byte
var tmp []byte
var dst []byte

func init() {
	text := []byte("maxwell backend is a newbest team,do u know ? yes, i know,i know")
	length := len(text)
	fmt.Println(length)
	bts = make([]byte, length)
	dst = make([]byte, length)
	tmp = make([]byte, length)
	copy(bts, []byte(text))

}

func Benchmark_self(b *testing.B) {
	for k := 0; k < b.N; k++ {
		Encrypt(bts[:], tmp[:])
		Decrypt(tmp[:], dst[:])
	}
}

func Benchmark_aes(b *testing.B) {
	for k := 0; k < b.N; k++ {
		AesEncrypt(bts[:], tmp[:])
		AesDecrypt(tmp[:], dst[:])
	}
}

func Benchmark_des(b *testing.B) {
	for k := 0; k < b.N; k++ {
		DesEncrypt(bts[:], tmp[:])
		DesDecrypt(tmp[:], dst[:])
	}
}

func Benchmark_rc4(b *testing.B) {
	for k := 0; k < b.N; k++ {
		Rc4Encrypt(bts[:], tmp[:])
		Rc4Decrypt(tmp[:], dst[:])
	}
}
