package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AdminIndexPage struct {
	Page
}

func adminIndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseFiles("templates/admin/index.html"))

	adminIndex := AdminIndexPage{Page: Page{Title: "Administration"}}

	tmpl.ExecuteTemplate(w, "layout", adminIndex)
}
