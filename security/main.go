package main

import (
	"fmt"
	"time"
)

func main() {
	var bts []byte
	var tmp []byte
	var dst []byte
	time := func(exe func(), log string) {
		start := time.Now()
		exe()
		end := time.Now()
		fmt.Println(log, len(bts), len(dst), end.Sub(start))
	}
	fill := func(text string, length int) {
		if length == 0 {
			length = len(text)
		}
		bts = make([]byte, length)
		if text != "" {
			copy(bts, []byte(text))
		} else {
			for i := range bts {
				bts[i] = byte(RandUint(0, 255))
			}
		}

		time(func() {
			tmp = Encrypt(bts)
		}, "self encrypt")

		time(func() {
			dst = Decrypt(tmp)
		}, "self decrypt")

		time(func() {
			tmp = AesEncrypt(bts)
		}, "aes encrypt")

		time(func() {
			dst = AesDecrypt(tmp)
		}, "aes decrypt")

		time(func() {
			tmp = DesEncrypt(bts)
		}, "des encrypt")

		time(func() {
			dst = DesDecrypt(tmp)
		}, "des decrypt")

		time(func() {
			tmp = Rc4Encrypt(bts)
		}, "rc4 encrypt")

		time(func() {
			dst = Rc4Decrypt(tmp)
		}, "rc4 decrypt")

		fmt.Println()
	}

	fill("this project is very good!", 0)
	fill("hello,this project is very good!", 0)
	fill("", 24)
	fill("", 255)
	fill("", 1025)
	fill("", 2045)
	fill("", 4045)
}
