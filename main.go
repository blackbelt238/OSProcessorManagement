package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2    // 2
var proc []*Processor // the processors
var ltime = 1         // time it take s to load a process is 1ms

// SchedulingProcedure defines the blueprint that all scheduling procedures must follow
type SchedulingProcedure interface {
	Schedule(que []*Job) int
	Name() string
}

type schedulingProcedure struct {
	fn   func(que [][2]int) int
	name string
}

// initProc initializes the slice representing the processors
func initProc() {
	proc = make([]*Processor, k)
	for i := 0; i < len(proc); i++ {
		proc[i] = CreateProcessor()
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
	sjn := &ShortestJobNext{}
	sim(rr)
	sim(sjn)

	b2 := []*Job{&Job{4, 9}, &Job{15, 2}, &Job{18, 16},
		&Job{20, 3}, &Job{26, 29}, &Job{29, 198},
		&Job{35, 7}, &Job{45, 170}, &Job{57, 180},
		&Job{83, 178}, &Job{88, 73}, &Job{95, 8}}
	fmt.Println(rr.Name(), "completed the b2 sequence in", rr.Schedule(b2), "ms")
	fmt.Println(sjn.Name(), "completed the b2 sequence in", sjn.Schedule(b2), "ms")
}
