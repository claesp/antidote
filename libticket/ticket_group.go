package libticket

// TicketGroup is a structure that represents a group within the ticket system.
type TicketGroup struct {
	ID      int
	Name    string
	Members []TicketUser
}
