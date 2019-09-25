package libticket

// TicketType indicates the type of the ticket.
type TicketType int

const (
	// TicketTypeIncident indicates that the ticket is a incident.
	TicketTypeIncident TicketType = iota
	// TicketTypeWorkorder indicates that the ticket is a workorder.
	TicketTypeWorkorder
	// TicketTypeProblem indicates that the ticket is a problem.
	TicketTypeProblem
	// TicketTypeChange indicates that the ticket is a change.
	TicketTypeChange
	// TicketTypeTask indicates that the ticket is a task.
	TicketTypeTask
)

func (tt TicketType) String() string {
	switch tt {
	case TicketTypeIncident:
		return "incident"
	case TicketTypeWorkorder:
		return "work order"
	case TicketTypeProblem:
		return "problem"
	case TicketTypeChange:
		return "change"
	case TicketTypeTask:
		return "task"
	default:
		return "unknown"
	}
}
