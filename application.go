package main;

import (
	"net/http"
)

type Application struct{
	middleware []http.Handler
	context Context
}

type CTXHandler struct{
	handler http.Handler
	ctx *Context
}

func (self *CTXHandler) ServeHTTP (w http.ResponseWriter,r *http.Request){
	self.handler.ServeHTTP(w,r);
}

func NewCTXHandler(f func(http.ResponseWriter,*http.Request)) *CTXHandler{
	ctxhandler := &CTXHandler{};
	handler := http.HandlerFunc(f);
	ctxhandler.handler = handler;
	return ctxhandler;
}

type Context struct{
	strs map[string]string
	ints map[string]int
	flts map[string]float32
}

func NewApplication() *Application{
	return &Application{};
}

func (self *Application) ServeHTTP (w http.ResponseWriter,r *http.Request){
	for _,handler := range self.middleware{
		handler.ServeHTTP(w,r);
	}
}

func (self *Application) Use (f func(http.ResponseWriter,*http.Request)){
	ctxhandler := NewCTXHandler(f);
	self.middleware = append(self.middleware,ctxhandler);
}