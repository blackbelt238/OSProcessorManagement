package main

// Processor represents how the processor will be simulated
type Processor struct {
	Job   *Job // the slot to hold the current job
	tproc int  // the time elapsed since the processor was last reset
}

// CreateProcessor returns a new processor
func CreateProcessor() *Processor {
	p := &Processor{}
	p.Reset()
	return p
}

// Process simulates the full processing of a job
func (p *Processor) Process() {
	p.tproc += p.Job.tproc
}

// LoadProcess loads a new job into the processor (taking load time into account)
func (p *Processor) LoadProcess(pr *Job) {
	p.Job = pr
	p.tproc += ltime
}

// Reset sets the processor's job to an invlaid job and resets tproc
func (p *Processor) Reset() {
	p.Job = CreateFillerJob()
	p.tproc = 0
}
