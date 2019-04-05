package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Page struct {
	Title string
}

type RootIndexPage struct {
	Page
}

func rootIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseFiles("templates/root/index.html"))

	rootIndex := RootIndexPage{Page: Page{Title: "Home"}}

	tmpl.ExecuteTemplate(w, "layout", rootIndex)
}

func main() {
	log.Println("antidote v.0.1")

	r := httprouter.New()
	r.ServeFiles("/css/*filepath", http.Dir("lib/css"))

	r.GET("/", rootIndexHandler)
	r.GET("/u/:id/open/", userOpenHandler)
	r.GET("/g/:id/open/", groupOpenHandler)
	r.GET("/t/:id/", ticketViewHandler)
	r.GET("/q/", searchIndexHandler)
	r.GET("/a/", adminIndexHandler)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
