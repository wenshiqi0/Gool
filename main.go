package main;

import (
	"net/http"
	"log"
	"fmt"
)

func YourHandler(ctx *Context) {
	ctx.set("test","Context test\n");
	ctx.res.Write([]byte("Good!Cool!\n"));
}

func Hello(ctx *Context){
	test := ctx.get("test");
	ctx.res.Write([]byte("Hello world\n"+test.(string)));
	fmt.Println(ctx.req.URL.Path);
}

func main() {
    app := NewApplication();
    app.Use(YourHandler);
	app.Use(Hello);
	log.Fatal(http.ListenAndServe(":3000", app));
}
