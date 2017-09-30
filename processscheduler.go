package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2 // 2
var proc = make([]int, k)

// circular simulates an on-line round-robin process scheduling procedure for a given job queue
// que holds information as [1][2,3] where:
//   1 is the queue position
//   2 is the arrival time
//   3 is the processing time required to finish
func circular(que [][2]int) int {
	tstart := que[0][0]
	tcur := tstart

	// initialize the processors
	p1 := que[0]
	p2 := que[1]
	que = que[2:]
	j := 1 // the processor upon which a process is run

	for i := 0; i < len(que); {
		// determine the processor to use
		j = (j + 1) % k

		// update time spent working on current process

		// move the current process to the end of the queue if it is incomplete
		//   do not increment list index

		// begin working on next process

	}

	return tcur - tstart // return the time elapsed since start
}

// processscheduler simulates a multi-core, non-preemptive process scheduler
func main() {
	fmt.Println(len(proc))
}
