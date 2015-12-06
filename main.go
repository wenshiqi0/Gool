package main;

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
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

func AddQuestion(ctx *Context,key string){
	ctx.req.ParseForm();
	result, err:= ioutil.ReadAll(ctx.req.Body);
	if(err != nil){
		ctx.req.Body.Close();
		log.Fatal(err);
		return;
	}
	ctx.req.Body.Close();
	var f interface{};
	json.Unmarshal([]byte(result), &f);
	
	m := f.(map[string]interface{});
	q := Question{};
		
	for k,v := range m {
		value := v.(string);
		switch k {
			case "title":
				q.Title = value;
			case "description":
				q.Description = value;
			case "topic":
				q.Topic = value;
		}
	}
	
	q.Key = key;
	
	mongodbQ := createMongodb();
	mongodbQ.ConnectMgo();
	mongodbQ.GetCollection("test","Question");
	mongodbQ.InsertQuestion(&q);
	mongodbQ.CloseMgo();	
	
	fmt.Println(q);
}

func FindQuestion(ctx *Context,key string){
	mongodbQ := createMongodb();
	mongodbQ.ConnectMgo();
	mongodbQ.GetCollection("test","Question");
	result := mongodbQ.FindQuestionByKey(key);
	mongodbQ.CloseMgo();
	
	fmt.Println(*result);
	
	body,err := json.Marshal(result);
	if(err != nil){
		log.Fatal(err);
	}
		
	ctx.res.Write(body);
}

func main() {
    app := NewApplication();
	
	app.Use(Logger);
	
	app.Use(NewRoute().
			Match("/index.html").
			Method("GET",Hello).
			Do);	
			
	app.Use(NewRouter().
			setUrl("/question").
			setGET(FindQuestion).
			setPOST(AddQuestion).
			Do);
		
	log.Fatal(http.ListenAndServe(":8000", app));	
}
