package libticket

type TicketPriority int

const (
	TicketPriorityLow TicketPriority = iota
	TicketPriorityMedium
	TicketPriorityHigh
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
