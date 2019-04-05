package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UserOpenPage struct {
	Page
	HasTickets bool
	Tickets    []TicketView
}

func userOpenHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("templates/user/open.html"))

	ticket_created, tc_err := time.Parse("2006-01-02 15:04", "2019-03-28 12:42")
	if tc_err != nil {
		ticket_created = time.Now()
	}
	ticket_ends, te_err := time.Parse("2006-01-02 15:04", "2019-03-30 19:55")
	if te_err != nil {
		ticket_ends = time.Now()
	}
	worklog1_time, w1t_err := time.Parse("2006-01-02 15:04", "2019-03-28 14:34")
	if w1t_err != nil {
		worklog1_time = time.Now()
	}
	ticket := Ticket{
		Number:      "INC0000556231",
		Title:       "Firewall changes",
		Description: "",
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
				Worklog: "Will start investigating this.",
				Created: worklog1_time,
				User: TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se"}}},
		Category: "Firewalls",
		Service:  "Data Communications",
		Status:   TicketStatusInProgress,
		Priority: TicketPriorityUrgent,
		Impact:   TicketImpactHigh,
		Type:     TicketTypeIncident,
		Assignee: TicketAssignee{
			Group: TicketGroup{
				ID:   3,
				Name: "Managed Services"}}}

	tickets := []TicketView{TicketView{ticket}}
	userOpen := UserOpenPage{Page: Page{Title: "Open"}, Tickets: tickets}
	userOpen.HasTickets = len(tickets) > 0

	tmpl.ExecuteTemplate(w, "layout", userOpen)
}
