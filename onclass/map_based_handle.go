package main

import (
	"net/http"
	"sync"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// 以method和urlPath组合字符串做key
	// handles map[string]func(ctx *Context)
	handles sync.Map
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(ctx *Context)) {

	key := h.key(method, pattern)
	// 根据HandlerBasedOnMap中的key匹配对应的处理函数xxx(ctx *Context)
	h.handles.Store(key, handlerFunc)
}

// 实现Handler接口的ServeHTTP方法
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.RequestURI)
	if handle, ok := h.handles.Load(key); ok {
		// 从map中取出处理函数xxx(ctx *Context)进行处理
		handle(NewCtx(c.W, c.R))
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{} // 确保该结构体实现了该接口，提醒程序员

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		// handles: make(map[string]func(ctx *Context), 4),
		handles: sync.Map{},
	}
}
