package libticket

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

type TicketView struct {
	Ticket
}

func (tv TicketView) CreatedAgeLabel() string {
	if tv.Ticket.CreatedAge() < (1 * time.Minute) {
		return "moments ago"
	}

	return humanize.Time(tv.Ticket.Created)
}

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

func (tv TicketView) EndsAtShortLabel() string {
	if tv.Ticket.EndsAt() > 0 {
		s := fmt.Sprintf("%s", humanize.RelTime(time.Now(), tv.Ticket.Ends, "", ""))
		return s[:len(s)-1]
	} else {
		return humanize.Time(tv.Ticket.Ends)
	}
}

func (tv TicketView) PriorityLabel() string {
	if tv.Priority == TicketPriorityUrgent {
		return "urgent"
	}

	return fmt.Sprintf("%s urgency", tv.Priority)
}

func (tv TicketView) IsUrgent() bool {
	return tv.Priority == TicketPriorityUrgent
}
