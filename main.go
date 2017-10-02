package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2    // 2
var proc []*Processor // the processors
var ltime = 1         // time it take s to load a process is 1ms

type SchedulingProcedure interface {
	Schedule(que []*Job) int
	Name() string
}

// initProc initializes the slice representing the processors
func initProc() {
	proc = make([]*Processor, k)
	for i := 0; i < len(proc); i++ {
		proc[i] = &Processor{}
		proc[i].Reset()
	}
}

// takes in 2 numbers and returns the maximum
func maxTimeElapsed() int {
	max := 0
	for i := 0; i < len(proc); i++ {
		if max < proc[i].tproc {
			max = proc[i].tproc
		}
	}
	return max
}

// processscheduler simulates a multi-core, non-preemptive process scheduler
func main() {
	rr := &RoundRobin{}
	e1 := []*Job{&Job{0, 3}, &Job{1, 2}, &Job{2, 1}}
	fmt.Println(rr.Schedule(e1))

	//sim(roundrobin)
	// b1 := genJobSeq()
	// b2 := // TODO: hardcode in the b2 dataset
}
