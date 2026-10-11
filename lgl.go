package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch a := a.(type) {
	case string:
		return strings.Repeat(a, b)
	case int:
		return a * b
	default:
		return a
	}
}

func main() {
	fmt.Println("result: ", Mul(2, 3))
	fmt.Println("result: ", Mul("b", 3))
}
