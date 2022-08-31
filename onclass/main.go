package main

import "net/http"

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}
	ctx.OKJson("ok")
}

func main() {
	server := NewHttpServer("my-server", MetricsFilterBuilder)
	server.Route(http.MethodGet, "/sign", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
