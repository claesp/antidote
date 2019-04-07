package main

import (
	"./templates"
	"github.com/claesp/antidote/libticket"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func groupOpen(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.GroupOpenPage{Tickets: []libticket.TicketView{}}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
