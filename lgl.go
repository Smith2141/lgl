package main

import (
	"fmt"
	"strings"
	"time"
)

type Stopwatch struct {
	startTime float64
	intervals []float64
}

func (sw *Stopwatch) Start() {
	sw.startTime = float64(time.Now().UnixNano())
}

func (sw *Stopwatch) SaveSplit() {
	newInterval := float64(time.Now().UnixNano())
	var interval float64 = (newInterval - sw.startTime)/ 1_000_000_000
	sw.intervals = append(sw.intervals, float64(interval))
}

func (sw *Stopwatch) GetResults() string {
	// return fmt.Printf("sw.intervals: %v\n", sw.intervals)
	var result strings.Builder
	for _, val := range sw.intervals {
		fmt.Fprintf(&result, "%.9fs ", val/1000000000)
	}

	result2 := result.String()
	result2 = strings.Trim(result2, " ")
	result2 = "[" + result2 + "]"
	return result2
}

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
