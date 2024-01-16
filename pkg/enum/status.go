package enum

const (
	Open = iota
	InProgress
	Completed
	Delayed
	Closed
	Suspended
)

type Status uint8

func (s Status) ToString() string {
	switch s {
	case Open:
		return "open"
	case InProgress:
		return "in progress"
	case Completed:
		return "completed"
	case Delayed:
		return "delayed"
	case Closed:
		return "closed"
	case Suspended:
		return "suspended"
	default:
		return "unknown"
	}
}
