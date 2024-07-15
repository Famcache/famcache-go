package jobs

import (
	"net"

	"github.com/famcache/famcache-go/domain"
	"github.com/google/uuid"
)

func (j *JobsManager) Registry(job domain.Job) domain.Job {
	uuid := uuid.New().String()

	command := j.commands.RegisterJob(uuid, job.Delay(), job.IsPeriodic())

	(*j.socket).Write([]byte(command))

	j.jobs[uuid] = job

	return job
}

func (j *JobsManager) Cancel(jobID string) error {
	delete(j.jobs, jobID)

	return nil
}

func (j *JobsManager) SetSocket(socket *net.Conn) {
	j.socket = socket
}

func (j *JobsManager) Trigger(reply domain.JobReply) {
	event := reply.Event()

	if event.IsRegistered() {
		job := j.jobs[reply.RequestID()]
		delete(j.jobs, reply.RequestID())

		j.jobs[reply.JobID()] = job

		job.(*Job).setJobID(reply.JobID())
	}

	if event.IsExecute() {
		job := j.jobs[reply.JobID()]

		// listenChan chan struct {}
		job.(*Job).listenChan <- struct{}{}
	}
}
