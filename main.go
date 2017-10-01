package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2 // 2
var proc [][2]int  // the processors
var ltime = 1      // time it take s to load a process is 1ms

// initProc initializes the slice representing the processors
func initProc() {
	proc = make([][2]int, k)
}

// takes in 2 numbers and returns the maximum
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// processscheduler simulates a multi-core, non-preemptive process scheduler
func main() {
	e1 := [][2]int{{0, 3}, {1, 2}, {2, 1}}
	fmt.Println(circular(e1))
	// b1 := genJobSeq()
	// b2 := // TODO: hardcode in the b2 dataset
}
