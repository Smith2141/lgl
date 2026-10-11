package main

import (
	// "fmt"
	"io"
	"log"
	"os"
	"strings"
)

func LimitReader(r io.Reader, n int) io.Reader {
	buf := make([]byte, n)
	var builder strings.Builder

	r.Read(buf)
	builder.Write(buf)
	limReader := strings.NewReader(builder.String())

	return limReader
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
