package step

// Action defines primitive action constituting a job.Job.
type Action int

const (
	GoToLocation Action = iota
	Pickup
	DropOff
	CollectCache
	CollectVoucher
	CallPhone
)

// Step is a part of Job execution
// JobID refers to job.Job step belongs to.
type Step struct {
	ID     string
	JobID  string
	Action Action
}