package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type ICipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

// AES
func NewAesCipher(key []byte) *AesCipher {
	var key32 [32]byte
	copy(key32[:], key)
	block, _ := aes.NewCipher(key32[:])
	return &AesCipher{key32: key32, block: block}
}

type AesCipher struct {
	key32 [32]byte
	block cipher.Block
}

func (this *AesCipher) Encrypt(src []byte) []byte {
	blockSize := this.block.BlockSize()
	src = PKCS5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(this.block, this.key32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

func (this *AesCipher) Decrypt(src []byte) []byte {
	blockSize := this.block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(this.block, this.key32[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = PKCS5UnPadding(dst)
	return dst
}

// DES
func NewDesCipher(key []byte) *DesCipher {
	var key8 [8]byte
	copy(key8[:], key)
	block, _ := des.NewCipher(key8[:])
	return &DesCipher{key8: key8, block: block}
}

type DesCipher struct {
	key8  [8]byte
	block cipher.Block
}

func (this *DesCipher) Encrypt(src []byte) []byte {
	blockSize := this.block.BlockSize()
	src = PKCS5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(this.block, this.key8[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

func (this *DesCipher) Decrypt(src []byte) []byte {
	blockSize := this.block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(this.block, this.key8[:blockSize])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = PKCS5UnPadding(dst)
	return dst
}

// RC4
func NewRc4Cipher(key []byte) *Rc4Cipher {
	var key32 [32]byte
	copy(key32[:], key)
	e, _ := rc4.NewCipher(key32[:])
	d, _ := rc4.NewCipher(key32[:])
	return &Rc4Cipher{e: e, d: d}
}

type Rc4Cipher struct {
	e *rc4.Cipher
	d *rc4.Cipher
}

func (this *Rc4Cipher) Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	this.e.XORKeyStream(dst, src)
	return dst
}

func (this *Rc4Cipher) Decrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	this.d.XORKeyStream(dst, src)
	return dst
}

// RSA
// pub 公钥文件路径
// priv 私钥文件路径
func NewRsaCipher(pub, priv string) *RsaCipher {
	data, _ := ioutil.ReadFile(pub)
	block, _ := pem.Decode(data)
	key, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pubk := key.(*rsa.PublicKey)

	data, _ = ioutil.ReadFile(priv)
	block, _ = pem.Decode(data)
	privk, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return &RsaCipher{
		pubk:  pubk,
		privk: privk,
	}
}

type RsaCipher struct {
	pubk  *rsa.PublicKey
	privk *rsa.PrivateKey
}

func (this *RsaCipher) Encrypt(src []byte) []byte {
	dst, _ := rsa.EncryptOAEP(sha1.New(), rand.Reader, this.pubk, src, []byte(""))
	return dst
}

func (this *RsaCipher) Decrypt(src []byte) []byte {
	dst, _ := rsa.DecryptOAEP(sha1.New(), rand.Reader, this.privk, src, []byte(""))
	return dst
}
