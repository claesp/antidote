package libticket

import (
	"time"

	humanize "github.com/dustin/go-humanize"
)

type TicketWorklog struct {
	ID      int
	Worklog string
	User    TicketUser
	Created time.Time
}

func (tw TicketWorklog) CreatedAge() time.Duration {
	return time.Since(tw.Created)
}

func (tw TicketWorklog) CreatedAgeLabel() string {
	return humanize.Time(tw.Created)
}
