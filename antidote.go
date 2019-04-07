package main

import (
	"log"

	"./templates"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func cssAntidote(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	ctx.SetContentType("text/css")
	templates.WriteCssAntidote(ctx)
}

func index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	ip := &templates.IndexPage{}
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

	log.Fatalln(fasthttp.ListenAndServe(":8080", router.Handler))
}
