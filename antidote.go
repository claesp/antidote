package main

import (
	"log"

	"./templates"
	"github.com/claesp/antidote/libticket"
	"github.com/claesp/antidote/libticket/drivers"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

var (
	TicketDB libticket.TicketDB = libticket.TicketDB{}
)

func cssAntidote(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	ctx.SetContentType("text/css")
	templates.WriteCssAntidote(ctx)
}

func index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	u := libticket.TicketUser{
		ID:   1,
		Name: "Claes Persson",
		Groups: []libticket.TicketGroup{
			libticket.TicketGroup{
				ID:   1,
				Name: "Nätdrift operation (Stockholm)"},
			libticket.TicketGroup{
				ID:   2,
				Name: "Nätdrift utveckling (Uppsala)"}}}
	ip := &templates.IndexPage{User: u}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, ip)
}

func main() {
	log.Println("antidote v.0.1")

	router := fasthttprouter.New()
	router.GET("/", index)
	router.GET("/css/antidote.css", cssAntidote)
	router.GET("/u/:id/open/", userOpen)
	router.GET("/g/:id/open/", groupOpen)
	router.GET("/t/:id/", ticketView)
	router.GET("/q/", searchIndex)
	router.GET("/a/", adminIndex)
	router.GET("/a/customer/new/", adminCustomerNew)
	router.GET("/a/group/new/", adminGroupNew)
	router.GET("/a/person/new/", adminPersonNew)
	router.GET("/a/user/new/", adminUserNew)

	TicketDB.Register("boltdb", &drivers.TicketDriverBoltDB{})
	conn_err := TicketDB.Connect()
	if conn_err != nil {
		log.Fatalln(conn_err)
	}

	log.Fatalln(fasthttp.ListenAndServe(":8080", router.Handler))
}
