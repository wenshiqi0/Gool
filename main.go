package main;

import (
	"net/http"
	"log"
)

func YourHandler(w http.ResponseWriter,r *http.Request,ctx *Context) {
	ctx.set("test","Context test\n");
	w.Write([]byte("Good!Cool!\n"));
}

func Hello(w http.ResponseWriter,r *http.Request,ctx *Context){
	test := ctx.get("test");
	w.Write([]byte("Hello world\n"+test.(string)));
}

func main() {
    app := NewApplication();
    app.Use(YourHandler);
	app.Use(Hello);
	log.Fatal(http.ListenAndServe(":3000", app));
}
