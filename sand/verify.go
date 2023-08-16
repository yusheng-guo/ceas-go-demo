package sand

import (
	"crypto"
	"crypto/rsa"
	"fmt"
)

// Verify 验签
func Verify(signData string, sign []byte, pk *rsa.PublicKey) error {
	hash := crypto.SHA1
	if !hash.Available() {
		return fmt.Errorf("crypto: requested hash function (%s) is unavailable", hash.String())
	}

	h := hash.New()
	h.Write([]byte(signData))

	return rsa.VerifyPKCS1v15(pk, hash, h.Sum(nil), sign)
}
