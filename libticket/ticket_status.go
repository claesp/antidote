package libticket

type TicketStatus int

const (
	TicketStatusNew TicketStatus = iota
	TicketStatusInProgress
	TicketStatusPending
	TicketStatusCompleted
	TicketStatusRejected
)

func (ts TicketStatus) String() string {
	switch ts {
	case TicketStatusNew:
		return "new"
	case TicketStatusInProgress:
		return "in progress"
	case TicketStatusPending:
		return "pending"
	case TicketStatusCompleted:
		return "completed"
	case TicketStatusRejected:
		return "rejected"
	default:
		return "unknown"
	}
}

func (ts TicketStatus) IsOpen() bool {
	switch ts {
	case TicketStatusNew:
		fallthrough
	case TicketStatusInProgress:
		return true
	default:
		return false
	}
}

func (ts TicketStatus) IsPaused() bool {
	switch ts {
	case TicketStatusPending:
		return true
	default:
		return false
	}
}
