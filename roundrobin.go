package main

// define
type RoundRobin struct{}

// implements SchedulingProcedure.Schedule
func (rr *RoundRobin) Schedule(que []*Job) int {
	return circular(que)
}

// implements SchedulingProcedure.Name
func (rr *RoundRobin) Name() string {
	return "Round Robin"
}

func circular(que []*Job) int {
	j := 0

	// initialize processors
	initProc()

	for i := 0; i < len(que); i++ {
		j = (j + 1) % k // determine the processor to use

		// as long as the processors already have processes in them,
		//   complete the current process
		if !proc[j].Job.IsFiller() {
			proc[j].Elapse(proc[j].Job.tproc)
		}
		proc[j].LoadProcess(que[i])
	}

	return maxTimeElapsed()
}
