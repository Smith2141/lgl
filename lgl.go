package main

import (
	"fmt"
	"github.com/Smith2141/lgl/example/randbyte"
	"github.com/Smith2141/lgl/hashbyte"
	"time"
)

func main() {

	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время — при каждом запуске оно будет разным

	buf := make([]byte, 16)

	for range 5 {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)

		hasher := hashbyte.New(0)
		hasher.Write(buf)
		fmt.Printf("Hash: %v \n", hasher.Hash())
	}

}
