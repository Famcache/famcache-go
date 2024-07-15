package jobs

func (j *Job) RequestID() string {
	return j.requestID
}

func (j *Job) Name() string {
	return j.name
}

func (j *Job) JobID() string {
	return j.jobID
}

func (j *Job) Delay() uint64 {
	return j.delay
}

func (j *Job) IsPeriodic() bool {
	return j.periodic
}

func (j *Job) Listen() <-chan struct{} {
	return j.listenChan
}

func (j *Job) IsRegistered() bool {
	return j.jobID != ""
}

func (j *Job) setJobID(jobID string) {
	j.jobID = jobID
}
