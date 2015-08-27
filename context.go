package main;

import (
	"net/http"
)

type Context struct{
	val map[interface{}]interface{}
	res http.ResponseWriter
	req *http.Request
}

func NewContext(res http.ResponseWriter,req *http.Request) *Context{
	ctx := &Context{};
	ctx.val = make(map[interface{}]interface{});
	ctx.res = res;
	ctx.req = req;
	return ctx;
}

func (self *Context) set(k interface{},v interface{}){
	self.val[k] = v;
}

func (self *Context) get(k interface{})interface{}{
	return self.val[k];
}
