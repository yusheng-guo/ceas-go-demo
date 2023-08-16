package utils

import (
	"crypto/rand"
	"fmt"
)

// RandomBytes 生成指定长度的随机字符串
func RandomBytes(length int) ([]byte, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return nil, fmt.Errorf("randomly generate string, err: %w", err)
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return bytes, nil
}
