package job

import (
	"dat320/lab4/scheduler/system/systime"
	"time"
)

func (j *Job) Scheduled(s systime.SystemTime) {
	j.SystemTime = s
	j.arrival = j.SystemTime.Now()
}

func (j *Job) Started(cpuID int) {
	if j.start == NotStartedYet {
		j.start = j.SystemTime.Now()
	}

}

func (j Job) TurnaroundTime() time.Duration {
	turnaroundTime := j.finished - j.arrival
	return turnaroundTime
}

func (j Job) ResponseTime() time.Duration {
	responseTime := j.start - j.arrival
	return responseTime
}
