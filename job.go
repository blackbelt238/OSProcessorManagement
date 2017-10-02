package main

type Job struct {
	tarr  int // time this job arrived at the queue
	tproc int // time this job requires to be fully processed
}

func (j *Job) IsFiller() bool {
	return j.tarr == -1 && j.tproc == -1
}

func FillerJob() *Job {
	return &Job{-1, -1}
}
