package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime int64
	marks     []float64
	Start
}

// func (sw Stopwatch) Start() {
	
// }

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
