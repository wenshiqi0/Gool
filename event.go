package gool;

/***********************************
Use the channel and goroutine to make
a Event struct to do the same thing as
Event in NodeJS.

Provide function Once() On() Dispatcher() 
and Emit() just like the thing in NodeJS

event := NewEvent();
Use this to get a event instance

event.On("test","test",test);
Use this to make a new event on "test".the
first param is the id for event.The second
param is the only param you can pass it to
the event function as the third param.It is
a interface{},so you can pass everything but
only one param.

event.Emit("test");
Use this to do the event function,and the event
function return will pass to this function to 
return by a channel

event.Once("test","test",test);
This is just like the On() function.The 
diffence is that this event can only be
emit once. 
************************************/

import (
	"fmt"
)

type Event struct{
	start map[string]chan string
	end map[string]chan interface{}
	done map[string]bool;
}

func NewEvent () *Event{
	event := &Event{};
	event.start = map[string]chan string{};
	event.end = map[string]chan interface{}{};
	event.done = map[string]bool{};
	return event;
}

func (self *Event)Once(reg string,param interface{},f func(interface{})interface{}){
	start := make(chan string);
	end := make(chan interface{});
	self.start[reg] = start;
	self.end[reg] = end;
	go func(){
		<-start;
		result := f(param);
		end<-result;
		delete(self.start,reg);
		delete(self.end,reg);
	}()
}

func (self *Event)On(reg string,param interface{},f func(interface{})interface{}){
	start := make(chan string);
	end := make(chan interface{});
	self.start[reg] = start;
	self.end[reg] = end;
	self.done[reg] = false;
	go func(){
		eventStart:
		{
			<-start;
			if(self.done[reg]){
				delete(self.done,reg);
				return;
			}
			result := f(param);
			end<-result;
		}
		goto eventStart;
	}()
}

func (self *Event)Dispatcher(reg string){
	start := self.start[reg];
	if(start == nil){
		return;
	}
	self.done[reg] = true;
	close(start);
	delete(self.start,reg);
	delete(self.end,reg);
}

func (self *Event)Emit(reg string) interface{} {
	start := self.start[reg];
	end := self.end[reg];
	if(start == nil){
		return nil;
	}
	start<-"result";
	result:=<-end
	return result;
}

func test(str interface{})interface{}{
	fmt.Println(str);
	result := "result";
	return result;
}

func main(){
	event := NewEvent();
	event.Once("haha","here",test);
	fmt.Println("gogo");
	result := event.Emit("haha");
	fmt.Println(result);
}