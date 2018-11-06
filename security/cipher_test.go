package security

import (
	"testing"
)

func Test_rsa(t *testing.T) {
	tmp = rsaCipher.Encrypt(bts)
	dst = rsaCipher.Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_des(t *testing.T) {
	tmp = desCipher.Encrypt(bts)
	dst = desCipher.Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_aes(t *testing.T) {
	tmp = aesCipher.Encrypt(bts)
	dst = aesCipher.Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}

func Test_rc4(t *testing.T) {
	tmp = rc4Cipher.Encrypt(bts)
	dst = rc4Cipher.Decrypt(tmp)
	t.Logf(string(dst))
	if string(bts) != string(dst) {
		t.Error("no eq")
	}
}
