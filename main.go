package main;

import (
	"net/http"
	"fmt"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"));
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello\n"));
	fmt.Println(r.URL.Path);
}

func main() {
    app := NewApplication();
	app.Use(Hello);
	app.Use(YourHandler);
    http.ListenAndServe(":3000", app);
}