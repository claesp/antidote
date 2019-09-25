package libticket

import (
	"time"
)

// Ticket is a structure representing a ticket.
type Ticket struct {
	Number      string
	Title       string
	Description string
	Contact     TicketContact
	Customer    TicketCustomer
	Created     time.Time
	Ends        time.Time
	Worklogs    []TicketWorklog
	Category    string
	Service     string
	Status      TicketStatus
	Priority    TicketPriority
	Impact      TicketImpact
	Type        TicketType
	Assignee    TicketAssignee
}

// CreatedAge calculates the age of the ticket.
func (t Ticket) CreatedAge() time.Duration {
	return time.Since(t.Created)
}

// EndsAt calculates the time until the ticket ends, generally when the SLA
// expires.
func (t Ticket) EndsAt() time.Duration {
	return time.Until(t.Ends)
}
