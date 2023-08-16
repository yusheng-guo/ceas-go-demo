package utils

import (
	"fmt"
	"testing"
)

func TestGenerateOrderNo(t *testing.T) {
	orderNo := GenerateOrderNo()
	fmt.Println(orderNo)
	// 202307201123021976800001
	// 202307201126487963160001
}
