package reply

import (
	"errors"
	"strings"

	"github.com/famcache/famcache-go/domain"
)

type JobEvent string

func (e JobEvent) IsExecute() bool {
	return e == JobEventTriggered
}

func (e JobEvent) IsRegistered() bool {
	return e == JobEventRegistered
}

func (e JobEvent) String() string {
	return string(e)
}

const (
	JobEventRegistered JobEvent = "JOB_EVENT_REGISTERED"
	JobEventTriggered  JobEvent = "JOB_EXECUTE"
)

type JobMessage struct {
	requestID string
	jobID     string
	event     domain.JobEvent
}

func (j JobMessage) JobID() string {
	return j.jobID
}

func (j JobMessage) RequestID() string {
	return j.requestID
}

func (j JobMessage) Event() domain.JobEvent {
	return j.event
}

func detectJobEvent(event string) (domain.JobEvent, error) {
	switch event {
	case "JOB":
		return JobEventRegistered, nil
	case "JOB_EXECUTE":
		return JobEventTriggered, nil
	default:
		return nil, errors.New("unknown job event")
	}
}

// <request_id> JOB OK <job_id>
// JOB_EXECUTE <job_id>
func NewJobReply(reply string) domain.JobReply {
	parts := normalizeReply(reply)

	if len(parts) != 4 {
		return JobMessage{
			requestID: "",
			jobID:     parts[1],
			event:     JobEventTriggered,
		}
	}

	return JobMessage{
		requestID: parts[0],
		jobID:     parts[3],
		event:     JobEventRegistered,
	}
}

func IsJobReply(reply string) bool {
	parts := strings.Split(strings.ReplaceAll(reply, "\n", ""), " ")

	return parts[1] == JOB_REPLY || parts[0] == JobEventTriggered.String()
}
