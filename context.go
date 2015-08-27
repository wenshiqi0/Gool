package main;

import(
	
)

type Context struct{
	val map[interface{}]interface{}
}

func (self *Context) set(k interface{},v interface{}){
	self.val[k] = v;
}

func (self *Context) get(k interface{})interface{}{
	return self.val[k];
}

func NewContext() *Context{
	ctx := &Context{};
	ctx.val = make(map[interface{}]interface{});
	return ctx;
}
