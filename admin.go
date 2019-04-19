package main

import (
	"./templates"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func adminCustomerNew(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminCustomerNewPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}

func adminGroupNew(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminGroupNewPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}

func adminIndex(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminIndexPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}

func adminPersonNew(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminPersonNewPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}

func adminUserCreate(ctx *fasthttp.RequestCtx, p fasthttprouter.Params) {

}

func adminUserNew(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	p := &templates.AdminUserNewPage{}
	ctx.SetContentType("text/html;charset=utf-8")
	templates.WritePageTemplate(ctx, p)
}
