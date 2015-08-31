package main;

import (
	"net/http"
	"log"
)

/*
func YourHandler(ctx *Context) {
	ctx.set("test","Context test\n");
	ctx.res.Write([]byte("Good!Cool!\n"));
}

func Hello(ctx *Context){
	test := ctx.get("test");
	header := ctx.res.Header();
	ctx.res.Write([]byte("Hello world\n"+test.(string)+"\n"+header.Get("Date")));
	fmt.Println(ctx.req.URL.Path);
}
*/

func Hello(ctx *Context,id string){
	ctx.res.Write([]byte(id));
}

func main() {
    app := NewApplication();
	
	app.Use(Logger)
	app.Use(NewRoute().Match("/index.html").Method("GET",Hello).Do);	
	
	log.Fatal(http.ListenAndServe(":3000", app));	
}
