package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
)

type TicketContact struct {
	ID    int
	Email string
	Name  string
	Phone string
}

func (tc TicketContact) HasEmail() bool {
	return len(tc.Email) > 0
}

func (tc TicketContact) HasPhone() bool {
	return len(tc.Phone) > 0
}

type TicketCustomer struct {
	ID   int
	Name string
}

type TicketUser struct {
	ID     int
	Name   string
	Email  string
	Groups []TicketGroup
}

func (tu TicketUser) PrimaryGroup() TicketGroup {
	return tu.Groups[0]
}

type TicketGroup struct {
	ID      int
	Name    string
	Members []TicketUser
}

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

type TicketAssignee struct {
	User  TicketUser
	Group TicketGroup
}

type TicketStatus int

const (
	TicketStatusNew TicketStatus = iota
	TicketStatusInProgress
	TicketStatusPending
	TicketStatusCompleted
	TicketStatusRejected
)

func (ts TicketStatus) String() string {
	switch ts {
	case TicketStatusNew:
		return "new"
	case TicketStatusInProgress:
		return "in progress"
	case TicketStatusPending:
		return "pending"
	case TicketStatusCompleted:
		return "completed"
	case TicketStatusRejected:
		return "rejected"
	default:
		return "unknown"
	}
}

func (ts TicketStatus) IsOpen() bool {
	switch ts {
	case TicketStatusNew:
		fallthrough
	case TicketStatusInProgress:
		return true
	default:
		return false
	}
}

func (ts TicketStatus) IsPaused() bool {
	switch ts {
	case TicketStatusPending:
		return true
	default:
		return false
	}
}

type TicketType int

const (
	TicketTypeIncident TicketType = iota
	TicketTypeWorkorder
	TicketTypeProblem
	TicketTypeChange
	TicketTypeTask
)

func (tt TicketType) String() string {
	switch tt {
	case TicketTypeIncident:
		return "incident"
	case TicketTypeWorkorder:
		return "work order"
	case TicketTypeProblem:
		return "problem"
	case TicketTypeChange:
		return "change"
	case TicketTypeTask:
		return "task"
	default:
		return "unknown"
	}
}

type TicketPriority int

const (
	TicketPriorityLow TicketPriority = iota
	TicketPriorityMedium
	TicketPriorityHigh
	TicketPriorityUrgent
)

func (tp TicketPriority) String() string {
	switch tp {
	case TicketPriorityLow:
		return "low"
	case TicketPriorityMedium:
		return "medium"
	case TicketPriorityHigh:
		return "high"
	case TicketPriorityUrgent:
		return "urgent"
	default:
		return "unknown"
	}
}

type TicketImpact int

const (
	TicketImpactLow TicketImpact = iota
	TicketImpactMedium
	TicketImpactHigh
)

func (ti TicketImpact) String() string {
	switch ti {
	case TicketImpactLow:
		return "low"
	case TicketImpactMedium:
		return "medium"
	case TicketImpactHigh:
		return "high"
	default:
		return "unknown"
	}
}

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

type TicketViewPage struct {
	Page
	Ticket TicketView
}

func ticketViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseFiles("templates/ticket/view.html"))

	id := ps.ByName("id")
	loc, loc_err := time.LoadLocation("Europe/Stockholm")
	if loc_err != nil {
		panic(loc_err)
	}
	ticket_created, tc_err := time.ParseInLocation("2006-01-02 15:04", "2019-02-28 12:42", loc)
	if tc_err != nil {
		ticket_created = time.Now()
	}
	ticket_ends, te_err := time.ParseInLocation("2006-01-02 15:04", "2019-03-30 20:06", loc)
	if te_err != nil {
		ticket_ends = time.Now()
	}
	worklog1_time, w1t_err := time.ParseInLocation("2006-01-02 15:04", "2019-03-28 14:34", loc)
	if w1t_err != nil {
		worklog1_time = time.Now()
	}
	worklog2_time, w2t_err := time.ParseInLocation("2006-01-02 15:04", "2019-03-30 02:15", loc)
	if w2t_err != nil {
		worklog2_time = time.Now()
	}
	ticket := Ticket{
		Number:      id,
		Title:       "Firewall changes",
		Description: "We need help with our firewall configuration. We need to open access for traffic between the 83.25.152.12 and 94.235.13.52 endpoints, on the following ports: tcp/80, tcp/443. Thanks in advance for your help!",
		Contact: TicketContact{
			ID:    1,
			Name:  "Anders Andersson",
			Email: "anders@example.com",
			Phone: "+46850650212"},
		Customer: TicketCustomer{
			ID:   2,
			Name: "Demoföretaget AB"},
		Created: ticket_created,
		Ends:    ticket_ends,
		Worklogs: []TicketWorklog{
			TicketWorklog{
				ID:      1,
				Worklog: "Will start investigating this. Reports back later.",
				Created: worklog1_time,
				User: TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se",
					Groups: []TicketGroup{
						TicketGroup{
							ID:   1,
							Name: "Nätdrift operations (Stockholm)"}}}},
			TicketWorklog{
				ID:      2,
				Worklog: "Having some issues with the configuration. Will try and contact end-user.",
				Created: worklog2_time,
				User: TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se",
					Groups: []TicketGroup{
						TicketGroup{
							ID:   1,
							Name: "Nätdrift operations (Stockholm)"}}}}},
		Category: "Communications",
		Service:  "Managed Firewalls",
		Status:   TicketStatusInProgress,
		Priority: TicketPriorityMedium,
		Impact:   TicketImpactLow,
		Type:     TicketTypeChange,
		Assignee: TicketAssignee{
			Group: TicketGroup{
				ID:   3,
				Name: "Managed Services"}}}

	tv := TicketView{Ticket: ticket}
	ticketView := TicketViewPage{Page: Page{Title: id}, Ticket: tv}

	tmpl.ExecuteTemplate(w, "layout", ticketView)
}
