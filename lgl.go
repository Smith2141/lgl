package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
// создаём strings.Builder
w := strings.Builder{}

for range 50 {
    // функция fmt.Fprintf принимает аргументом io.Writer
    // благодаря этому можно записывать форматированный вывод
    fmt.Fprintf(&w, "%v", math.NaN())
}

w.Write([]byte("... BATMAN!"))

// выводим собранную строку
fmt.Printf("%s\n", &w)
}
