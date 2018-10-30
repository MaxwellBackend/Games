package main

import (
	"testing"
)

var bts []byte
var tmp []byte
var dst []byte

func init() {
	text := []byte("maxwell backend is a newbest team,do u know ? yes, i know,i know")
	length := len(text)
	bts = make([]byte, length)
	copy(bts, []byte(text))

}

func Benchmark_self(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = Encrypt(bts)
		dst = Decrypt(tmp)
	}
}

func Benchmark_aes(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = AesEncrypt(bts)
		dst = AesDecrypt(tmp)
	}
}

func Benchmark_des(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = DesEncrypt(bts)
		dst = DesDecrypt(tmp)
	}
}

func Benchmark_rc4(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = Rc4Encrypt(bts)
		dst = Rc4Decrypt(tmp)
	}
}
