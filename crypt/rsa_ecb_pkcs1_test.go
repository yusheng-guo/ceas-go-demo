package crypt

import (
	"fmt"
	"testing"
)

func TestRSAEncryptECB(t *testing.T) {
	ecb, err := RSAEncryptECB([]byte("123456"), "../cert/sand_public.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println(ecb)
}
