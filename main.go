package main;

import (
	"net/http"
	"log"
)

func Hello(ctx *Context,id string){
	ctx.res.Write([]byte(id));
}

func RegisterUser(ctx *Context,name string){
	ctx.req.ParseForm();
	user := ctx.req.PostForm.Get("pass");
	ctx.res.Write([]byte(name));
	ctx.res.Write([]byte(user));
}

func GetUser(ctx *Context,name string){
	ctx.res.Write([]byte(name));
	ctx.res.Write([]byte("this is a GET method"));
}

func main() {
    app := NewApplication();
	
	app.Use(Logger);
	
	app.Use(NewRoute().
			Match("/index.html").
			Method("GET",Hello).
			Do);	
					
	app.Use(NewRouter().
			setUrl("/user").
			setGET(GetUser).
			setPOST(RegisterUser).
			Do);
	
	log.Fatal(http.ListenAndServe(":3000", app));	
}
