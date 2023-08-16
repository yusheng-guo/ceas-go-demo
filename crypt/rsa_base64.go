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
	encryptKeyBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, key)
	if err != nil {
		return "", fmt.Errorf("aes encrypt, err: %w", err)
	}
	// 2.3 Base64编码
	return base64.StdEncoding.EncodeToString(encryptKeyBytes), nil
}
