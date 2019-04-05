package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SearchIndexPage struct {
	Page
	HasTickets bool
	Tickets    []Ticket
}

func searchIndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseFiles("templates/search/index.html"))

	tickets := []Ticket{}

	searchIndex := SearchIndexPage{Page: Page{Title: "Search"}, Tickets: tickets}
	searchIndex.HasTickets = len(tickets) > 0

	tmpl.ExecuteTemplate(w, "layout", searchIndex)
}
