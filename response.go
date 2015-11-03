package main;

import (
	"net/http"
)

type Response struct{
	w http.ResponseWriter
	statusCode string
	body []byte
	method string
}

func (self *Response)Write(){
	
}