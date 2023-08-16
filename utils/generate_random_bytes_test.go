package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	str, err := RandomBytes(16)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
