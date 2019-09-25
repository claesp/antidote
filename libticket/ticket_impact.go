package libticket

// TicketImpact is the impact ticket has.
type TicketImpact int

const (
	// TicketImpactLow indicates low impact.
	TicketImpactLow TicketImpact = iota
	// TicketImpactMedium indicates medium impact.
	TicketImpactMedium
	// TicketImpactHigh indicates high impact.
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
