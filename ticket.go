package main

import (
	"time"

	"github.com/claesp/antidote/templates"

	"github.com/claesp/antidote/libticket"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func ticketView(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	id := ps.ByName("id")
	loc, loc_err := time.LoadLocation("Europe/Stockholm")
	if loc_err != nil {
		panic(loc_err)
	}
	ticket_created, tc_err := time.ParseInLocation("2006-01-02 15:04", "2019-04-06 12:42", loc)
	if tc_err != nil {
		ticket_created = time.Now()
	}
	ticket_ends, te_err := time.ParseInLocation("2006-01-02 15:04", "2019-04-08 12:42", loc)
	if te_err != nil {
		ticket_ends = time.Now()
	}
	worklog1_time, w1t_err := time.ParseInLocation("2006-01-02 15:04", "2019-04-07 05:34", loc)
	if w1t_err != nil {
		worklog1_time = time.Now()
	}
	worklog2_time, w2t_err := time.ParseInLocation("2006-01-02 15:04", "2019-04-07 02:15", loc)
	if w2t_err != nil {
		worklog2_time = time.Now()
	}
	ticket := libticket.Ticket{
		Number:      id,
		Title:       "Firewall changes",
		Description: "We need help with our firewall configuration. We need to open access for traffic between the 83.25.152.12 and 94.235.13.52 endpoints, on the following ports: tcp/80, tcp/443. Thanks in advance for your help!",
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
				Worklog: "Will start investigating this. Reports back later.",
				Created: worklog1_time,
				User: libticket.TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se",
					Groups: []libticket.TicketGroup{
						libticket.TicketGroup{
							ID:   1,
							Name: "Nätdrift operations (Stockholm)"}}}},
			libticket.TicketWorklog{
				ID:      2,
				Worklog: "Having some issues with the configuration. Will try and contact end-user.",
				Created: worklog2_time,
				User: libticket.TicketUser{
					ID:    1,
					Name:  "Claes Persson",
					Email: "claes.persson@ip-only.se",
					Groups: []libticket.TicketGroup{
						libticket.TicketGroup{
							ID:   1,
							Name: "Nätdrift operations (Stockholm)"}}}}},
		Category: "Communications",
		Service:  "Managed Firewalls",
		Status:   libticket.TicketStatusInProgress,
		Priority: libticket.TicketPriorityUrgent,
		Impact:   libticket.TicketImpactLow,
		Type:     libticket.TicketTypeChange,
		Assignee: libticket.TicketAssignee{
			Group: libticket.TicketGroup{
				ID:   3,
				Name: "Managed Services"}}}

	tv := libticket.TicketView{Ticket: ticket}
	p := &templates.TicketViewPage{Ticket: tv}
	ctx.SetContentType("text/html; charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
