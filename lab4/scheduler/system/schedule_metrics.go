package system

import (
	"dat320/lab4/scheduler/job"
	"math"
	"time"
)

// Avg returns the average of a metric defined by the function f.
func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration {
	sum := time.Duration(0)
	var counter float64 = 0
	for _, j := range sch {
		sum += f(j.Job)
		counter += 1
	}
	regneut := math.Round(float64(sum) / counter)
	finalSum := time.Duration(regneut)
	return finalSum
}

func (sch Schedule) AvgResponseTime() time.Duration {
	sum := time.Duration(0)
	var counter float64 = 0
	for _, j := range sch {
		sum += j.Job.ResponseTime()
		counter += 1
	}
	regneut := float64(sum) / counter
	finalSum := time.Duration(regneut)
	return finalSum
}

func (sch Schedule) AvgTurnaroundTime() time.Duration {
	sum := time.Duration(0)
	var counter float64 = 0
	for _, j := range sch {
		sum += j.Job.TurnaroundTime()
		counter += 1
	}
	regneut := float64(sum) / counter
	finalSum := time.Duration(regneut)
	return finalSum
}
