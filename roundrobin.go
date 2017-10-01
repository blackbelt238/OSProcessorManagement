package main

// define
type RoundRobin struct{}

// implements SchedulingProcedure.Schedule
func (rr *RoundRobin) Schedule(que [][2]int) int {
	return circular(que)
}

// implements SchedulingProcedure.Name
func (rr *RoundRobin) Name() string {
	return "Round Robin"
}

func circular(que [][2]int) int {
	j := 0

	// loads a process into a processor
	loadProcess := func(loadfrom int) {
		proc[j].Job = que[loadfrom]
	}

	// initialize processors
	initProc()

	for i := 0; i < len(que); i++ {
		j = (j + 1) % k // determine the processor to use

		// as long as the processors already have processes in them
		if k < i {

		}
	}

	return 0
}
