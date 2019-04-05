package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type GroupOpenPage struct {
	Page
	HasTickets bool
	Tickets    []Ticket
}

func groupOpenHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseFiles("templates/group/open.html"))

	tickets := []Ticket{}

	groupOpen := GroupOpenPage{Page: Page{Title: "Open"}, Tickets: tickets}
	groupOpen.HasTickets = len(tickets) > 0

	tmpl.ExecuteTemplate(w, "layout", groupOpen)
}
