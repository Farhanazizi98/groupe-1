package rr

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type roundRobin struct {
	queue   job.Jobs
	cpu     *cpu.CPU
	quantum time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *roundRobin {
	return &roundRobin{
		queue:   make(job.Jobs, 0),
		cpu:     cpus[0],
		quantum: quantum,
	}
}

func (rr *roundRobin) Add(job *job.Job) {
	rr.queue = append(rr.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.

func (rr *roundRobin) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	slice := systemTime%rr.quantum == 0
	if rr.cpu.IsRunning() {
		if rr.cpu.Tick() {
			jobsFinished++
		}
	}
	if slice {
		rr.reassign()
	}
	return jobsFinished
}

// reassign assigns a job to the cpu

func (rr *roundRobin) reassign() {
	currentJob := rr.cpu.CurrentJob()
	if currentJob != nil {
		rr.Add(currentJob)
	}
	nxtJob := rr.getNewJob()
	rr.cpu.Assign(nxtJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job

func (rr *roundRobin) getNewJob() *job.Job {
	if len(rr.queue) == 0 {
		return nil
	}
	removedJob := rr.queue[0]
	rr.queue = rr.queue[1:]
	return removedJob
}
