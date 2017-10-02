package main

type Processor struct {
	Job   *Job // the slot to hold the current job
	tproc int  // the time elapsed since the processor was last reset
}

func (p *Processor) Elapse(t int) {
	p.tproc += t
}

func (p *Processor) LoadProcess(pr *Job) {
	p.Job = pr
	p.tproc += ltime
}

func (p *Processor) Reset() {
	p.Job = FillerJob()
	p.tproc = 0
}

func (p *Processor) ProcessingTime() int {
	return p.tproc
}
