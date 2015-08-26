package main;

import (
	"net/http"
)

type Application struct{
	middleware []func(http.ResponseWriter,*http.Request,*Context)
	context *Context
}

func NewApplication() *Application{
	app := &Application{};
	app.context = NewContext();
	return app;
}

func (self *Application) ServeHTTP (w http.ResponseWriter,r *http.Request){
	for _,f := range self.middleware{
		f(w,r,self.context)
	}
}

func (self *Application) Use (f func(http.ResponseWriter,*http.Request,*Context)){
	self.middleware = append(self.middleware,f);
}


