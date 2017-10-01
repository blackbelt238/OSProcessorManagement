package main

import (
	"fmt"
	"math"
	"math/rand"
)

func genJobs() (randJobs [][2]int) {
	for i := 0; i < 1000; i++ {
		tarr := rand.Intn(500) + 1
		tproc := rand.Intn(500) + 1
		randJobs = append(randJobs, [2]int{tarr, tproc})
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

func sim(sf *schedulingProcedure) {
	res := [1000]int{}

	tot := 0
	max := -1
	min := math.MaxInt16

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		res[i] = sf.fn(genJobs())

		tot += res[i]
		if res[i] > max {
			max = res[i]
		}
		if min > res[i] {
			min = res[i]
		}
	}

	fmt.Println(sf.name, " achieved the following on 1000 simulations:\n\tavg:", tot/1000, "\n\tsd:", sd(res, tot/1000), "\ntmin:", min, "\n\tmax:", max)
}
