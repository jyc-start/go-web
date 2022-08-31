package main

import (
	"net/http"
)

type Server interface {
	Routable
	Start(address string) error
}

// sdkHttpServer 基于http库实现
type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

// Route 注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	// v1写法
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	// 工厂+闭包
	//	ctx := NewCtx(writer, request)
	//	handlerFunc(ctx)
	//})
	// v2写法
	//key := s.handler.key(method, pattern)
	//s.handler.handles[key] = handlerFunc

	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := NewCtx(w, r)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	// 因为是一个链，所以把最后的业务逻辑处理也作为一环
	var root Filter = handler.ServeHTTP

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
