package main

// RoundRobin is a scheduling procedure using the round-robin method of scheduling
type RoundRobin struct{}

// Schedule implements SchedulingProcedure.Schedule
func (rr *RoundRobin) Schedule(que []*Job) int {
	j := 0

	// initialize processors
	initProc()

	for i := 0; i < len(que); i++ {
		j = (j + 1) % k // determine the processor to use

		// as long as the processor already has a real job, complete it
		if !proc[j].Job.IsFiller() {
			proc[j].Process()
			proc[j].Wait(que[i].tarr - proc[j].Job.tarr - proc[j].Job.tproc)
		}
		proc[j].LoadProcess(que[i])
	}

	return maxTimeElapsed() + que[0].tarr
}

// Name implements SchedulingProcedure.Name
func (rr *RoundRobin) Name() string {
	return "Round Robin"
}
