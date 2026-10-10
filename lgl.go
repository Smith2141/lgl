package main

import (
	"fmt"
	"time"

	"github.com/Smith2141/lgl/example/randbyte"
)

func main() {

	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

	// buf := make([]byte, 16)
	buf := make([]byte, 24)

	for range 5 {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

}
