package ecrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"unsafe"
)

func AESEncryptECB(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// ECBEncrypter is an implementation of the cipher.BlockMode interface
// that uses Electronic Codebook (ECB) mode.
type ECBEncrypter struct {
	b         cipher.Block
	blockSize int
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
func (x *ECBEncrypter) CryptBlocks(dst, src []byte) {
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
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// InexactOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
func InexactOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 && (uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1])))
}
