package crypt

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func AESEncryptECB(pt, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return base64.StdEncoding.EncodeToString(ct), nil
}

//// ECBEncrypter is an implementation of the cipher.BlockMode interface
//// that uses Electronic Codebook (ECB) mode.
//type ECBEncrypter struct {
//	b         cipher.Block
//	blockSize int
//}
//
//func AESEncryptECB(plaintext []byte, key []byte) (string, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//
//	blockSize := block.BlockSize()
//	plaintext = PKCS5Padding(plaintext, blockSize)
//
//	ciphertext := make([]byte, len(plaintext))
//	ecb := NewECBEncrypter(block)
//	ecb.CryptBlocks(ciphertext, plaintext)
//
//	return base64.StdEncoding.EncodeToString(ciphertext), nil
//}
//
//// PKCS5Padding PKCS5填充
//func PKCS5Padding(data []byte, blockSize int) []byte {
//	padding := blockSize - (len(data) % blockSize)
//	padText := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(data, padText...)
//}
//
//// NewECBEncrypter create a new ECBEncrypter.
//func NewECBEncrypter(b cipher.Block) *ECBEncrypter {
//	return &ECBEncrypter{
//		b:         b,
//		blockSize: b.BlockSize(),
//	}
//}
//
//// BlockSize returns the block size of the cipher.
//func (x *ECBEncrypter) BlockSize() int {
//	return x.blockSize
//}
//
//// CryptBlocks encrypts a full block.
//func (x *ECBEncrypter) CryptBlocks(dst, src []byte) {
//	if len(src)%x.blockSize != 0 {
//		panic("crypto/cipher: input not full blocks")
//	}
//	if len(dst) < len(src) {
//		panic("crypto/cipher: output smaller than input")
//	}
//	if InexactOverlap(dst[:len(src)], src) {
//		panic("crypto/cipher: invalid buffer overlap")
//	}
//
//	for len(src) > 0 {
//		x.b.Encrypt(dst, src[:x.blockSize])
//		src = src[x.blockSize:]
//		dst = dst[x.blockSize:]
//	}
//}
//
//// InexactOverlap reports whether x and y share memory at any (not necessarily
//// corresponding) index. The memory beyond the slice length is ignored.
//func InexactOverlap(x, y []byte) bool {
//	return len(x) > 0 && len(y) > 0 && (uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
//		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1])))
//}
