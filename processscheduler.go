package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2 // 2
var proc = make([]int, k)

// processscheduler simulates a multi-core, non-preemptive process scheduler
func main() {
	fmt.Println(len(proc))
}
