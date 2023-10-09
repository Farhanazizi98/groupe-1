package stride

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type stride struct {
	queue   job.Jobs
	cpu     *cpu.CPU
	quantum time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *stride {
	return &stride{
		queue:   make(job.Jobs, 0),
		cpu:     cpus[0],
		quantum: quantum,
	}
}

func (s *stride) Add(job *job.Job) {
	s.queue = append(s.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.

func (s *stride) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	sliceExpired := systemTime%s.quantum == 0
	if s.cpu.IsRunning() {
		currentJob := s.cpu.CurrentJob()
		currentJob.Pass += currentJob.Stride
		if s.cpu.Tick() {
			jobsFinished++
		}
	}
	if sliceExpired {
		s.reassign()
	}
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *stride) reassign() {
	currentJob := s.cpu.CurrentJob()
	if currentJob != nil {
		s.Add(currentJob)
	}
	nxtJob := s.getNewJob()
	s.cpu.Assign(nxtJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *stride) getNewJob() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}

	i := MinPass(s.queue)
	removedJob := s.queue[i]
	sliceBefore := s.queue[:i]
	sliceafter := s.queue[i+1:]
	s.queue = append(sliceBefore, sliceafter...)

	return removedJob

}

// minPass returns the index of the job with the lowest pass value.
func MinPass(theJobs job.Jobs) int {
	lowest := 0

	for i, j := range theJobs {
		if j.Pass < theJobs[lowest].Pass {
			lowest = i
		} else if j.Pass == theJobs[lowest].Pass {
			if j.Stride < theJobs[lowest].Stride {
				lowest = i
			}
		}
	}

	return lowest
}
