package libticket

type TicketGroup struct {
	ID      int
	Name    string
	Members []TicketUser
}
