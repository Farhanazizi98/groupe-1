package stride

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// NewJob creates a job for stride scheduling.
func NewJob(size, tickets int, estimated time.Duration) *job.Job {
	const numerator = 10_000
	job := job.New(size, estimated)
	if tickets < 1 {
		tickets = 1
	}
	job.Tickets = tickets
	job.Pass = 0
	job.Stride = (numerator / tickets)
	return job
}
