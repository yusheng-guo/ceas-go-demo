package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"unsafe"
)

// ECBEncrypter 是cipher.BlockMode接口的实现 使用Electronic Codebook (ECB)模式进行解密
type ECBEncrypter struct {
	b         cipher.Block
	blockSize int
}

// ECBDecrypter 是cipher.BlockMode接口的实现 使用Electronic Codebook (ECB)模式进行解密
type ECBDecrypter struct {
	b         cipher.Block
	blockSize int
}

// NewECBDecrypter 创建一个新的ECBDecrypter。
func NewECBDecrypter(b cipher.Block) *ECBDecrypter {
	return &ECBDecrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// BlockSize returns the block size of the cipher.
func (x *ECBEncrypter) BlockSize() int {
	return x.blockSize
}

// NewECBEncrypter create a new ECBEncrypter.
func NewECBEncrypter(b cipher.Block) *ECBEncrypter {
	return &ECBEncrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// BlockSize 返回密码的块大小。
func (x *ECBDecrypter) BlockSize() int {
	return x.blockSize
}

// AESEncryptECB AES加密 ECB模式
func AESEncryptECB(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecryptECB AES解密 ECB模式
func AESDecryptECB(ciphertext string, key []byte) ([]byte, error) {
	// 解码Base64密文
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	// 检查密文长度是否是块大小的倍数
	if len(decodedCiphertext)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: ciphertext length is not a multiple of the block size")
	}

	ecb := NewECBDecrypter(block)

	// 解密密文
	plaintext := make([]byte, len(decodedCiphertext))
	ecb.CryptBlocks(plaintext, decodedCiphertext)

	// 去除PKCS5填充
	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}

// PKCS5Padding PKCS5填充
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// CryptBlocks encrypts a full block.
func (x *ECBEncrypter) CryptBlocks(dst, src []byte) error {
	if len(src)%x.blockSize != 0 {
		return errors.New("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		return errors.New("crypto/cipher: output smaller than input")
	}
	if InexactOverlap(dst[:len(src)], src) {
		return errors.New("crypto/cipher: invalid buffer overlap")
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
	return nil
}

// InexactOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
func InexactOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 && (uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1])))
}

// CryptBlocks 解密一个完整块
func (x *ECBDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	if InexactOverlap(dst[:len(src)], src) {
		panic("crypto/cipher: invalid buffer overlap")
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// PKCS5UnPadding 移除PKCS5填充。
func PKCS5UnPadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}
