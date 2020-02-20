package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
)

var decrypted string

func init() {
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
}

func MainRsa() {
	var data []byte
	var err error
	if decrypted != "" {
		data, err = base64.StdEncoding.DecodeString(decrypted)
		if err != nil {
			panic(err)
		}
	} else {
		data, err = RsaEncrypt([]byte("origin data ...dxc"))
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`  
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDJTjjQR52ODTeToW7jbVnjiYLFuou98nu3JWlIjM4155RdvRld
Tsem/uRHy91yLmDh6xXoLFESqKgQRkUTI8aBHGruu892uQvA5FaWWxxLzrlV45He
8qhmT+f1THT7h/YcWkwW6K4ajI/z0xfGURb1tZKNAexPyjxosB7utayuLwIDAQAB
AoGBAJ2I6G8eTE85Sq/umo/CKKEnIL7KMMeyNlzqGw1am5mVpAcDkBD0MMwgFteU
07SiVNEXnuo1nDCE6hjqVO4YhiNqnTn5vw1IgUh14lOMQUZ7YQhpeIgaF7Cas6X5
TBFZ+4u4McS5N0/Ynz3I9iefnqPxdWdCC+HZNebZUEkXy3eJAkEA5VITO0cW+gTB
Yp23uKVBlMO5Y9PGuPp+oYMTDdOcW9OfUXKLRjzDgjMYDw6Rfla1NakMzQw2Ipr4
W115iVEKrQJBAOC5xBypiuFPrgOtaXc/0zi6m42W3QmqvKW7hzWpD3pB6p/pIR7g
SUJctOkofsigemaZME4YyhbsOfZe0sJD88sCQEB9Ckd7QGzi9XdEHyxf3Md3GyWF
orbWStIkyDD1N11jb8Q50AzafaiZscRaNnQu8hq7BYyMSJUGu50F9m21SqkCQCUS
H+rGgRV36nHK2noEeliCWAS8XUwp8SK5060jbV0yoHyunXsVbqBW4LURrBB2gJqK
LLDFYBj18P3WWJgZU48CQQDZGHShLXFQRvFCuDM8SFtTJhiCGsA3GS43RuCYWqNf
E5c6NB55hX3/ybMwVJtWke3EO8Fzdbgqrt7QT5FL1Z/z
-----END RSA PRIVATE KEY-----  
`)

var publicKey = []byte(`  
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJTjjQR52ODTeToW7jbVnjiYLF
uou98nu3JWlIjM4155RdvRldTsem/uRHy91yLmDh6xXoLFESqKgQRkUTI8aBHGru
u892uQvA5FaWWxxLzrlV45He8qhmT+f1THT7h/YcWkwW6K4ajI/z0xfGURb1tZKN
AexPyjxosB7utayuLwIDAQAB
-----END PUBLIC KEY-----  
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}


///-----------
/**
	加密
 */
func Rsa2Encrypt(plainText,publicKey[]byte)([]byte, error)  {
	block, _ := pem.Decode(publicKey)
	if block == nil{
		return  nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}


/**
	解密
 */
func Rsa2Decrypt(ciphertext,privateKey[]byte)([]byte, error)  {
	block, _ := pem.Decode(privateKey)
	if block == nil{
		return  nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil,err
	}
	return rsa.DecryptPKCS1v15(rand.Reader,priv,ciphertext)
}