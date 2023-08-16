package sand

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

// Sign 加签
func Sign(data []byte) (sign string, err error) {
	h := crypto.Hash.New(crypto.SHA1)
	// 对数据进行签名
	_, err = h.Write(data)
	if err != nil {
		return "", fmt.Errorf("write data, err: %w", err)
	}
	hashed := h.Sum(nil)
	// 加载私钥
	r, err := LoadPrivateKey("./cert/sand_private.pfx", "nft123456")
	if err != nil {
		return "", fmt.Errorf("load private key, err: %w", err)
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, r, crypto.SHA1, hashed)
	if err != nil {
		return "", fmt.Errorf("sign, err: %w", err)
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return sign, nil
}
