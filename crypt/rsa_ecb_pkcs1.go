package crypt

import (
	"ceas-go-demo/utils"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// RSAEncryptECB RSA加密并进行Base64编码
func RSAEncryptECB(key []byte, path string) (string, error) {
	// 1 加载公钥
	rsaPublicKey, err := utils.LoadPublicKey(path)
	if err != nil {
		return "", fmt.Errorf("load public key, err: %w", err)
	}
	// 2 公钥加密
	//encryptKeyBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, key)
	//if err != nil {
	//	return "", fmt.Errorf("aes encrypt, err: %w", err)
	//}
	encryptKeyBytes, err := encrypt(key, rsaPublicKey)
	if err != nil {
		return "", err
	}

	// 3 Base64编码
	return base64.StdEncoding.EncodeToString(encryptKeyBytes), nil
}

const RESERVE_SIZE = 11

// encrypt 分块加密
func encrypt(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	keySize := publicKey.Size()
	fmt.Println(keySize)
	encryptBlockSize := keySize - RESERVE_SIZE // 加密块大小

	//nBlock := len(plainText) / encryptBlockSize
	//if len(plainText)%encryptBlockSize != 0 {
	//	nBlock++
	//}

	encrypted := make([]byte, 0)

	for offset := 0; offset < len(plainText); offset += encryptBlockSize {
		end := offset + encryptBlockSize
		if end > len(plainText) {
			end = len(plainText)
		}
		fmt.Println(end - offset)
		block, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText[offset:end])
		if err != nil {
			return nil, err
		}
		encrypted = append(encrypted, block...)
	}

	return encrypted, nil
}
