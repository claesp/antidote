package libticket

import (
	"time"

	humanize "github.com/dustin/go-humanize"
)

// TicketWorklog is a representation of a single work log item within a ticket.
type TicketWorklog struct {
	ID      int
	Worklog string
	User    TicketUser
	Created time.Time
}

// CreatedAge calculate the age of the work log.
func (tw TicketWorklog) CreatedAge() time.Duration {
	return time.Since(tw.Created)
}

// CreatedAgeLabel shows a human representation of the work logs creation time.
func (tw TicketWorklog) CreatedAgeLabel() string {
	return humanize.Time(tw.Created)
}
