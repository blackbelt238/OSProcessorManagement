package main

// ShortestJobNext is a scheduling procedure using the shortest-first method of scheduling
type ShortestJobNext struct{}

// Schedule implements SchedulingProcedure.Schedule
func (sjn *ShortestJobNext) Schedule(que []*Job) int {
	initProc() // reset the processors
	tstart := que[0].tarr
	j := 0

	// loop until there are no processes left in que
	for len(que) > 0 {
		var jshori int // the index of the shortest job
		var jshor *Job

		j = (j + 1) % k // determine the processor to use

		// find the shortest job in the queue
		for i := 0; i < len(que); i++ {
			if jshor == nil || que[i].tproc < jshor.tproc {
				jshor = que[i]
				jshori = i
			}
		}

		// remove the job from the queue
		copy(que[jshori:], que[jshori+1:])
		que = que[:len(que)-1]

		// as long as the processor already has a real job, complete it
		if !proc[j].Job.IsFiller() {
			proc[j].Process()
			proc[j].Wait(jshor.tarr - proc[j].Job.tarr - proc[j].Job.tproc)
		}
		proc[j].LoadProcess(jshor)
	}

	return maxTimeElapsed() + tstart
}

// Name implements SchedulingProcedure.Name
func (sjn *ShortestJobNext) Name() string {
	return "Shortest Job Next"
}
