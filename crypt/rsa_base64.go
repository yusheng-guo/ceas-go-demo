package crypt

import (
	"ceas-go-demo/utils"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// RSAAndBase64 RSA加密并进行Base64编码
func RSAAndBase64(key []byte, path string) (string, error) {
	// 2.1 加载公钥
	rsaPublicKey, err := utils.LoadPublicKey(path)
	if err != nil {
		return "", fmt.Errorf("load public key, err: %w", err)
	}
	// 2.2 公钥加密
	//encryptKeyBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, key)
	//if err != nil {
	//	return "", fmt.Errorf("aes encrypt, err: %w", err)
	//}
	encryptKeyBytes, err := encrypt(key, rsaPublicKey)
	if err != nil {
		return "", err
	}

	// 2.3 Base64编码
	return base64.StdEncoding.EncodeToString(encryptKeyBytes), nil
}

func encrypt(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	// 计算加密块的大小
	keySize := publicKey.Size()
	encryptBlockSize := keySize - 11

	// 计算加密块的数量
	nBlock := len(plainText) / encryptBlockSize
	if len(plainText)%encryptBlockSize != 0 {
		nBlock++
	}

	encrypted := make([]byte, 0)

	// 分块加密
	for offset := 0; offset < len(plainText); offset += encryptBlockSize {
		end := offset + encryptBlockSize
		if end > len(plainText) {
			end = len(plainText)
		}

		// 执行加密操作
		block, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText[offset:end])
		if err != nil {
			return nil, err
		}

		encrypted = append(encrypted, block...)
	}

	return encrypted, nil
}
