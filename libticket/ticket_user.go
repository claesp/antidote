package libticket

// TicketUser represents a user in the ticket system.
type TicketUser struct {
	ID     int
	Name   string
	Email  string
	Groups []TicketGroup
}

// PrimaryGroup returns the primary group for the user (i.e. the first one)
func (tu TicketUser) PrimaryGroup() TicketGroup {
	return tu.Groups[0]
}
