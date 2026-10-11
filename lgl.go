package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch va := a.(type) {
	case string:
		return strings.Repeat(va, b)
	case int:
		return va * b
	case fmt.Stringer:
		return strings.Repeat(va.String(), b)
	default:
		return nil
	}
}

func main() {
	fmt.Println("result: ", Mul(2, 3))
	fmt.Println("result: ", Mul("b", 3))
}
