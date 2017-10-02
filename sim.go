package main

import (
	"fmt"
	"math"
	"math/rand"
)

func genJobs() (randJobs []*Job) {
	for i := 0; i < 1000; i++ {
		tarr := i
		tproc := rand.Intn(500) + 1
		randJobs = append(randJobs, &Job{tarr, tproc})
	}
	return
}

func sd(in [1000]int, mean int) float64 {
	ssq := 0
	for i := range in {
		ssq += int(math.Pow(float64(i-mean), 2))
	}
	v := ssq / 999
	return math.Sqrt(float64(v))
}

func sim(sf SchedulingProcedure) {
	res := [1000]int{}

	tot := 0
	max := -1
	min := math.MaxInt16

	for i := 0; i < 1000; i++ {
		res[i] = sf.Schedule(genJobs())

		tot += res[i]
		if res[i] > max {
			max = res[i]
		}
		if min > res[i] {
			min = res[i]
		}
	}

	fmt.Println(sf.Name(), "achieved the following statistics on 1000 simulations (all in ms):\n\tavg:", tot/1000, "\n\tsd:", sd(res, tot/1000), "\n\tmin:", min, "\n\tmax:", max)
}
