package sjf

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"sort"
	"time"
)

type sjf struct {
	queue job.Jobs
	cpu   *cpu.CPU
}

func New(cpus []*cpu.CPU) *sjf {
	return &sjf{
		queue: make(job.Jobs, 0),
		cpu:   cpus[0],
	}
}

func (s *sjf) Add(job *job.Job) {
	s.queue = append(s.queue, job)
	sort.Sort(s.queue)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.

func (s *sjf) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	if s.cpu.IsRunning() {
		if s.cpu.Tick() {
			jobsFinished++
			s.reassign()
		}
	} else {
		s.reassign()
	}
	return jobsFinished
}

// reassign assigns a job to the cpu

func (s *sjf) reassign() {
	nxtJob := s.getNewJob()
	s.cpu.Assign(nxtJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job

func (s *sjf) getNewJob() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}
	removedJob := s.queue[0]
	s.queue = s.queue[1:]
	return removedJob
}
