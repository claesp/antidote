package libticket

// TicketPriority marks the priority for the ticket.
type TicketPriority int

const (
	// TicketPriorityLow indicates low priority.
	TicketPriorityLow TicketPriority = iota
	// TicketPriorityMedium indicates medium priority.
	TicketPriorityMedium
	// TicketPriorityHigh indicates high priority.
	TicketPriorityHigh
	// TicketPriorityUrgent indicates urgent priority.
	TicketPriorityUrgent
)

func (tp TicketPriority) String() string {
	switch tp {
	case TicketPriorityLow:
		return "low"
	case TicketPriorityMedium:
		return "medium"
	case TicketPriorityHigh:
		return "high"
	case TicketPriorityUrgent:
		return "urgent"
	default:
		return "unknown"
	}
}
