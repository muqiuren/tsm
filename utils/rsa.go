package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"os"
)

const (
	PrivateKeyName = "tsm-private.pem"
	PublicKeyName  = "tsm-public.pem"
)

var (
	PriKey []byte
	PubKey []byte
)

func init() {
	if exists, _ := PathExists(BasePath + "/" + PrivateKeyName); !exists {
		err := genRsaKey(2048)
		CheckErr(err)
	}
	loadRsaKey()
}

// genRsaKey 生成rsa密钥
func genRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(BasePath + "/" + PrivateKeyName)
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(BasePath + "/" + PublicKeyName)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// loadRsaKey 加载rsa密钥
func loadRsaKey() {
	var err error
	PriKey, err = os.ReadFile(BasePath + "/" + PrivateKeyName)
	CheckErr(err)

	PubKey, err = os.ReadFile(BasePath + "/" + PublicKeyName)
	CheckErr(err)
}

// EncryptByPublic 公钥加密
func EncryptByPublic(plain string, publicKey []byte) (cipherByte []byte, err error) {
	msg := []byte(plain)
	pubBlock, _ := pem.Decode(publicKey)
	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	CheckErr(err)

	pub := pubKeyValue.(*rsa.PublicKey)
	encryptOAEP, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, msg, nil)
	CheckErr(err)
	cipherByte = encryptOAEP
	return
}

// DecryptByPrivate 私钥解密
func DecryptByPrivate(cipherByte []byte, privateKey []byte) (plainText string, err error) {
	priBlock, _ := pem.Decode(privateKey)
	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	CheckErr(err)

	decryptOAEP, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	CheckErr(err)

	plainText = string(decryptOAEP)
	return
}
