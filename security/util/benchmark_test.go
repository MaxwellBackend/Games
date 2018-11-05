package util

import (
	"testing"
)

func Benchmark_aes(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = aesCipher.Encrypt(bts)
		dst = aesCipher.Decrypt(tmp)
	}
}

func Benchmark_des(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = desCipher.Encrypt(bts)
		dst = desCipher.Decrypt(tmp)
	}
}

func Benchmark_rc4(b *testing.B) {
	for k := 0; k < b.N; k++ {
		tmp = rc4Cipher.Encrypt(bts)
		dst = rc4Cipher.Decrypt(tmp)
	}
}
