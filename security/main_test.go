package main

import (
	"testing"
)

func init() {
	text := []byte("maxwell backend is a newbest team,do u know ? yes, i know,i know")
	length := len(text)
	bts = make([]byte, length)
	copy(bts, []byte(text))
}

func Test_self(t *testing.T) {
	tmp = Encrypt(bts)
	dst = Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_des(t *testing.T) {
	tmp = DesEncrypt(bts)
	dst = DesDecrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_aes(t *testing.T) {
	tmp = AesEncrypt(bts)
	dst = AesDecrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_rc4(t *testing.T) {
	tmp = Rc4Encrypt(bts)
	dst = Rc4Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}
