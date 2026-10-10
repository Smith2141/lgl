package main

import (
	"fmt"
	"time"

	"github.com/Smith2141/lgl/example/randbyte"
)

func main() {

	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.
	// var num64 int64 = 1

	buf := make([]byte, 16)

	for range 5 {
		// generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.
		// generator := randbyte.New(num64) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

}
