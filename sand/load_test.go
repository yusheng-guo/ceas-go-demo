package sand

import (
	"fmt"
	"testing"
)

func TestLoadPublicKey(t *testing.T) {
	key, err := LoadPublicKey("../cert/sand_public.cer")
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
}

func TestLoadPrivateKey(t *testing.T) {
	key, err := LoadPrivateKey("../cert/sand_private_backup.pfx", "nft123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
}
