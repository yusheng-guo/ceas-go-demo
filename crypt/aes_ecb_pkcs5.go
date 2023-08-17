package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"unsafe"
)

// AESEncryptECB ECB模式AES加密 并Base64编码
//func AESEncryptECB(pt, key []byte) (string, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//	mode := ecb.NewECBEncrypter(block)
//	padder := padding.NewPkcs5Padding()
//	pt, err = padder.Pad(pt) // pad last block of plaintext if block size less than block cipher size
//	if err != nil {
//		return "", fmt.Errorf("pad, err: %w", err)
//	}
//	ct := make([]byte, len(pt))
//	mode.CryptBlocks(ct, pt)
//	return base64.StdEncoding.EncodeToString(ct), nil
//}

// AESDecryptECB ECB模式AES解密
//func AESDecryptECB(ct, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, fmt.Errorf("new cipher, err: %w", err)
//	}
//	mode := ecb.NewECBDecrypter(block)
//	pt := make([]byte, len(ct))
//	mode.CryptBlocks(pt, ct)
//	padder := padding.NewPkcs5Padding()
//	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
//	if err != nil {
//		return nil, fmt.Errorf("padder unpad, err: %w", err)
//	}
//	return pt, nil
//}

// ECBEncrypter is an implementation of the cipher.BlockMode interface
// that uses Electronic Codebook (ECB) mode.
type ECBEncrypter struct {
	b         cipher.Block
	blockSize int
}

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

// PKCS5Padding PKCS5填充
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// NewECBEncrypter create a new ECBEncrypter.
func NewECBEncrypter(b cipher.Block) *ECBEncrypter {
	return &ECBEncrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// BlockSize returns the block size of the cipher.
func (x *ECBEncrypter) BlockSize() int {
	return x.blockSize
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
