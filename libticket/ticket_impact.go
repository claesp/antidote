package libticket

type TicketImpact int

const (
	TicketImpactLow TicketImpact = iota
	TicketImpactMedium
	TicketImpactHigh
)

func (ti TicketImpact) String() string {
	switch ti {
	case TicketImpactLow:
		return "low"
	case TicketImpactMedium:
		return "medium"
	case TicketImpactHigh:
		return "high"
	default:
		return "unknown"
	}
}
