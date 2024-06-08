package main

import (
	"time"

	"github.com/claesp/antidote/templates"

	"github.com/claesp/antidote/libticket"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func userOpen(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	ticket_created, tc_err := time.Parse("2006-01-02 15:04", "2019-04-06 12:42")
	if tc_err != nil {
		ticket_created = time.Now()
	}
	ticket_ends, te_err := time.Parse("2006-01-02 15:04", "2019-04-08 12:42")
	if te_err != nil {
		ticket_ends = time.Now()
	}
	worklog1_time, w1t_err := time.Parse("2006-01-02 15:04", "2019-04-07 14:34")
	if w1t_err != nil {
		worklog1_time = time.Now()
	}
	ticket := libticket.Ticket{
		Number:      "INC0000556231",
		Title:       "Firewall changes",
		Description: "",
		Contact: libticket.TicketContact{
			ID:    1,
			Name:  "Anders Andersson",
			Email: "anders@example.com",
			Phone: "+46850650212"},
		Customer: libticket.TicketCustomer{
			ID:   2,
			Name: "Demoföretaget AB"},
		Created: ticket_created,
		Ends:    ticket_ends,
		Worklogs: []libticket.TicketWorklog{
			libticket.TicketWorklog{
				ID:      1,
				Worklog: "Will start investigating this.",
				Created: worklog1_time,
				User: libticket.TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se"}}},
		Category: "Firewalls",
		Service:  "Data Communications",
		Status:   libticket.TicketStatusInProgress,
		Priority: libticket.TicketPriorityUrgent,
		Impact:   libticket.TicketImpactHigh,
		Type:     libticket.TicketTypeIncident,
		Assignee: libticket.TicketAssignee{
			Group: libticket.TicketGroup{
				ID:   3,
				Name: "Managed Services"}}}
	ticket2 := libticket.Ticket{
		Number:      "INC000055823",
		Title:       "Can't connect to website",
		Description: "",
		Contact: libticket.TicketContact{
			ID:    1,
			Name:  "Anders Andersson",
			Email: "anders@example.com",
			Phone: "+46850650212"},
		Customer: libticket.TicketCustomer{
			ID:   2,
			Name: "ICA Gruppen AB"},
		Created: ticket_created,
		Ends:    ticket_ends,
		Worklogs: []libticket.TicketWorklog{
			libticket.TicketWorklog{
				ID:      1,
				Worklog: "Will start investigating this.",
				Created: worklog1_time,
				User: libticket.TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se"}}},
		Category: "Firewalls",
		Service:  "Data Communications",
		Status:   libticket.TicketStatusInProgress,
		Priority: libticket.TicketPriorityMedium,
		Impact:   libticket.TicketImpactMedium,
		Type:     libticket.TicketTypeIncident,
		Assignee: libticket.TicketAssignee{
			Group: libticket.TicketGroup{
				ID:   3,
				Name: "Managed Services"}}}

	tickets := []libticket.TicketView{libticket.TicketView{ticket}, libticket.TicketView{ticket2}}
	p := &templates.UserOpenPage{Tickets: tickets}
	ctx.SetContentType("text/html; charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
