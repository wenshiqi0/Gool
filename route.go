package main;

import (
	"regexp"
	"strings"
	//"fmt"
)

type Route struct{
	mapping string
	id string
	method string
	do func(*Context,string)
}

func NewRoute() *Route{
	route := &Route{};
	route.mapping = "";
	route.do = nil;
	return route;
}

func (self *Route)Match(url string) *Route{
	self.mapping,self.id = match(url);
	return self;
}

func (self *Route)Method(method string,f func(*Context,string)) *Route{
	self.method = method;
	self.do = f;
	return self;
}

func (self *Route)Do(ctx *Context) int{
	mapping,id := match(ctx.req.URL.Path);
	if(self.mapping == mapping && self.method == ctx.req.Method){
		self.do(ctx,id);
		return 0;
	}
	return 1;
}

func (self *Route)IsEmpty() bool{
	if(self.do == nil || self.mapping == ""){
		return true;
	}
	return false;
}

func match(str string)(mapping string,id string){
	reg , _ := regexp.Compile(`^[a-zA-Z0-9\_\-\/\.]*\:([a-zA-Z0-9\/\.\|]*)$`);
	array := reg.FindAllSubmatch([]byte(str),-1);
	if(array == nil){
		return str,"";
	}
	mapping = strings.Split(string(array[0][0]),":")[0];
	id = string(array[0][1]);
	return;
}