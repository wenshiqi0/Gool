package main;

import (
	"time"
	"fmt"
	"strings"
)

type LogInfo struct{
	url string
	method string
	time time.Time
}

func onLog(param interface {})interface{}{
	info := param.(*LogInfo);
	past := info.time;
	now := time.Now();
	inter := now.Sub(past);
	comeOut := strings.Join([]string{"<-- ",info.url," ",info.method," "},"");
	fmt.Print(comeOut);
	fmt.Println(inter);
	fmt.Print("\n");
	return inter;
}

func Logger(ctx *Context){
	comeTime := time.Now();
	url := ctx.req.URL.Path;
	method := ctx.req.Method;
	comeIn := strings.Join([]string{"--> ",url," ",method},"");
	fmt.Println(comeIn);
	info := &LogInfo{url,method,comeTime};
	ctx.event.Once("log",info,onLog);
}