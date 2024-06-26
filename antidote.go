//go:generate go install github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates

package main

import (
	"log"

	"github.com/claesp/antidote/templates"
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
	var u libticket.TicketUser
	user, err := TicketDB.GetUser(1)
	if err != nil {
		u = libticket.TicketUser{
			ID:   1,
			Name: "Claes Persson",
			Groups: []libticket.TicketGroup{
				libticket.TicketGroup{
					ID:   1,
					Name: "Active Network Services Operation (Stockholm)"},
				libticket.TicketGroup{
					ID:   2,
					Name: "Active Network Services Utveckling (Uppsala)"}}}
	} else {
		u = user
	}
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
	router.POST("/a/user/new/", adminUserCreate)

	TicketDB.Register("bbolt", &drivers.TicketDriverBbolt{})
	conn_err := TicketDB.Connect()
	if conn_err != nil {
		log.Fatalln(conn_err)
	}
	defer TicketDB.Disconnect()

	log.Println("listening on :8080")
	log.Fatalln(fasthttp.ListenAndServe(":8080", router.Handler))
}
