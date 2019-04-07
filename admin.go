package main

import (
	"./templates"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func adminIndex(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminIndexPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
