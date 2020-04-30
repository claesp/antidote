package libticket

// TicketStatus indicates the current status for the ticket.
type TicketStatus int

const (
	// TicketStatusNew indicates that the ticket is new.
	TicketStatusNew TicketStatus = iota
	// TicketStatusInProgress indicates that the ticket is currently in
	// progress.
	TicketStatusInProgress
	// TicketStatusPending indicates that the ticket is awaiting some external
	// information, and is currently paused.
	TicketStatusPending
	// TicketStatusCompleted indicates that the ticket has been resolved.
	TicketStatusCompleted
	// TicketStatusRejected indicates that the ticket has been rejected.
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

// IsOpen returns true if the status is in a state that could be considered
// open or active.
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

// IsPaused returns true if the ticket is in a paused state.
func (ts TicketStatus) IsPaused() bool {
	switch ts {
	case TicketStatusPending:
		return true
	default:
		return false
	}
}
