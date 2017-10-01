package main

import (
	"fmt"
)

// k represents the number of cores the processor has
var k = 5244%3 + 2 // 2
var proc [][2]int  // the processors
var ltime = 1      // time it take s to load a process is 1ms

// takes in 2 numbers and returns the maximum
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// circular simulates an on-line round-robin process scheduling procedure for a given job queue
// que holds information as [1][2,3] where:
//   1 is the queue position
//   2 is the arrival time
//   3 is the processing time required to finish
func circular(que [][2]int) int {
	// loads a process into a processor
	loadProcess := func(loadto int) {
		proc[loadto] = que[0]
		que = que[1:]
	}
	// finds the index of the first process that has an arrival time greater
	//   than that of the given time
	findIndex := func(arrTime int) int {
		i := 0
		for ; i < len(que); i++ {
			if que[i][0] > arrTime {
				break
			}
		}
		return i
	}

	tstart := que[0][0]
	ttot := tstart // total time elapsed since the first process was recieved

	// initialize the processors (load in the first k processes)
	proc = make([][2]int, k)
	for i := 0; i < k; i++ {
		loadProcess(i)
	}

	j := k - 1 // the processor upon which a process is run

	// while there are processes in the queue, swap and process
	for i := 0; i < len(que); {
		pnext := que[i] // find the next process
		j = (j + 1) % k // determine the processor to use

		// update time spent working on current process
		ttot = pnext[0]
		proc[j][1] -= func() int {
			tmp := ttot - proc[j][0] - ltime
			if tmp < 0 {
				tmp = 0
			}
			return tmp
		}()

		// re-insert the process into the queue if processing is still required
		if proc[j][1] > 0 {
			proc[j][0] = ttot + 1

			// re-insert at the proper place in the queue based on arrival time
			ins := findIndex(ttot)
			que = append(que, [2]int{0, 0})
			copy(que[ins+1:], que[ins:])
			que[ins] = proc[j]
		}

		// begin working on next process
		loadProcess(j)
	}

	// handle the last 2 processes
	ttot += max(proc[0][1], proc[1][1])

	return ttot - tstart // return the time elapsed since start
}

// processscheduler simulates a multi-core, non-preemptive process scheduler
func main() {
	e1 := [][2]int{{0, 3}, {1, 2}, {2, 1}}
	fmt.Println(circular(e1))
	// b1 := genJobSeq()
	// b2 := // TODO: hardcode in the b2 dataset
}
