package sand

import (
	"bytes"
	"ceas-go-demo/utils"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

// AESEncrypt AES加密
func AESEncrypt(data map[string]any) (string, string, error) {
	// 生成16字节的AES密钥
	aesKey, err := utils.RandomBytes(16)
	if err != nil {
		return "", "", err
	}

	// 对数据进行AES加密
	plainValue, err := json.Marshal(data)
	if err != nil {
		return "", "", err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", "", err
	}
	paddedValue := Pad(plainValue, block.BlockSize())
	iv := make([]byte, aes.BlockSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	encryptValueBytes := make([]byte, len(paddedValue))
	mode.CryptBlocks(encryptValueBytes, paddedValue)
	encryptValue := base64.StdEncoding.EncodeToString(encryptValueBytes)

	// 对AES Key 进行RSA加密
	rsaPublicKey, err := LoadPublicKey("./cert/sand_public.cer")
	if err != nil {
		return "", "", fmt.Errorf("load public key, err: %w", err)
	}
	encryptKeyBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, aesKey)
	if err != nil {
		return "", "", fmt.Errorf("aes encrypt, err: %w", err)
	}
	sandEncryptKey := base64.StdEncoding.EncodeToString(encryptKeyBytes)

	// 返回加密后的结果 包括value和key
	return encryptValue, sandEncryptKey, nil
}

// Pad 对数据进行填充
func Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
