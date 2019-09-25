package libticket

// TicketAssignee is a structure representing the current assignee for a ticket.
type TicketAssignee struct {
	User  TicketUser
	Group TicketGroup
}
