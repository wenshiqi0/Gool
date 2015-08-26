package main;

import (
	"net/http"
	"fmt"
)

func YourHandler(w http.ResponseWriter,r *http.Request,ctx *Context) {
	ctx.setStr("test","Context test\n");
	w.Write([]byte("Good!Cool!\n"));
}

func Hello(w http.ResponseWriter,r *http.Request,ctx *Context){
	test := ctx.getStr("test");
	fmt.Println(test);
	w.Write([]byte("Hello world\n"+test));
}

func main() {
    app := NewApplication();
    app.Use(YourHandler);
	app.Use(Hello);
    http.ListenAndServe(":3000", app);
}
