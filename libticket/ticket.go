package libticket

import (
	"time"
)

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

func (t Ticket) CreatedAge() time.Duration {
	return time.Since(t.Created)
}

func (t Ticket) EndsAt() time.Duration {
	return time.Until(t.Ends)
}
