package main;

import (
	"net/http"
)

type Application struct{
	middleware []func(*Context)
	context *Context
}

func NewApplication() *Application{
	app := &Application{};
	return app;
}

func (self *Application) ServeHTTP (w http.ResponseWriter,r *http.Request){
	self.context = NewContext(w,r);
	for _,f := range self.middleware{
		f(self.context)
	}
	self.context.event.Emit("log");
}

func (self *Application) Use (f func(*Context)){
	self.middleware = append(self.middleware,f);
}


