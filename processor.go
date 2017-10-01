package main

type Processor struct {
	Job   [2]int // the slot to hold the current job
	tproc int    // the time elapsed since the processor was last reset
}

func (p *Processor) Reset() {
	p.tproc = 0
}

func (p *Processor) ProcessingTime() int {
	return p.tproc
}
