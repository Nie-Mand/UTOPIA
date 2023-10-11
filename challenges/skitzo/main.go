package main

import (
	"fmt"
)

const (
	x = 11
	y = 13
	z = 23
)

func exp(base, exponent, modulus int64) int64 {
	if modulus == 1 {
		return 0
	}
	base = base % modulus
	result := int64(1)
	for i := int64(0); i < exponent; i++ {
		result = (result * base) % modulus
	}
	return result
}

func main() {
	key := "key_here"

	encrypted := ""
	for _, c := range key {
		c := exp(int64(c), z, x * y)
		encrypted += fmt.Sprintf("%02x", c)
	}

	// 143, 3d1e58653624280b40653624280b0b4065883d6c6c
	fmt.Println(x * y, encrypted)
}