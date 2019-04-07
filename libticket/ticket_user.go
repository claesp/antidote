package libticket

type TicketUser struct {
	ID     int
	Name   string
	Email  string
	Groups []TicketGroup
}

func (tu TicketUser) PrimaryGroup() TicketGroup {
	return tu.Groups[0]
}
