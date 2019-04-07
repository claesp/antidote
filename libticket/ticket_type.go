package libticket

type TicketType int

const (
	TicketTypeIncident TicketType = iota
	TicketTypeWorkorder
	TicketTypeProblem
	TicketTypeChange
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
