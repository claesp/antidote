package libticket

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// TicketView is a custom view for a Ticket.
type TicketView struct {
	Ticket
}

// CreatedAgeLabel returns a human representation of the creation time for the
// ticket.
func (tv TicketView) CreatedAgeLabel() string {
	if tv.Ticket.CreatedAge() < (1 * time.Minute) {
		return "moments ago"
	}

	return humanize.Time(tv.Ticket.Created)
}

// EndsAtLabel returns a human representation for the end of the ticket,
// generally when the SLA expires.
func (tv TicketView) EndsAtLabel() string {
	if tv.Ticket.EndsAt() > 0 {
		if tv.Ticket.EndsAt() < (1 * time.Minute) {
			return "expires shortly"
		} else if tv.Ticket.EndsAt() == 0 {
			return "expired now"
		} else {
			s := fmt.Sprintf("expires in %s", humanize.RelTime(time.Now(), tv.Ticket.Ends, "", ""))
			return s[:len(s)-1]
		}
	} else {
		return fmt.Sprintf("expired %s", humanize.Time(tv.Ticket.Ends))
	}
}

// EndsAtShortLabel is a human representation for the end of the ticket,
// generally when the SLA expires. But unlike EndsAtLabel() this is a shorter
// form.
func (tv TicketView) EndsAtShortLabel() string {
	if tv.Ticket.EndsAt() > 0 {
		s := fmt.Sprintf("%s", humanize.RelTime(time.Now(), tv.Ticket.Ends, "", ""))
		return s[:len(s)-1]
	}
	return humanize.Time(tv.Ticket.Ends)
}

// PriorityLabel returns a human representation of the priority of the ticket.
func (tv TicketView) PriorityLabel() string {
	if tv.Priority == TicketPriorityUrgent {
		return "urgent"
	}

	return fmt.Sprintf("%s urgency", tv.Priority)
}

// IsUrgent returns true if the ticket is marked as urgent.
func (tv TicketView) IsUrgent() bool {
	return tv.Priority == TicketPriorityUrgent
}
