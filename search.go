package main

import (
	"github.com/claesp/antidote/templates"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func searchIndex(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.SearchIndexPage{Query: ""}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
