package main;

import(
	
)

type Context struct{
	ints map[string]int
	strs map[string]string
	flts map[string]float32
}

func NewContext() *Context{
	ctx := &Context{};
	ctx.ints = make(map[string]int);
	ctx.strs = make(map[string]string);
	ctx.flts = make(map[string]float32);
	return ctx;
}

func (self *Context) setInt(k string,v int) {
	self.ints[k] = v;
}

func (self *Context) getInt(k string) int{
	return self.ints[k];
}

func (self *Context) setStr(k string,v string) {
	self.strs[k] = v;
}

func (self *Context) getStr(k string) string{
	return self.strs[k];
}

func (self *Context) setFlt(k string,v float32) {
	self.flts[k] = v;
}

func (self *Context) getFlt(k string) float32{
	return self.flts[k];
}


