package main;

type Router struct{
	get *Route
	post *Route
	put *Route
	del *Route
	
	url string
}

func NewRouter() *Router{
	router := &Router{};
	return router;
}

func (self *Router)setUrl(u string) *Router{
	self.url = u;
	self.get = NewRoute().Match(self.url);
	self.post = NewRoute().Match(self.url);
	self.put = NewRoute().Match(self.url);
	self.del = NewRoute().Match(self.url);
	return self;
}

func (self *Router)setGET(f func(*Context,string)) *Router{
	self.get = self.get.Method("GET",f);
	return self;
}

func (self *Router)setPOST(f func(*Context,string)) *Router{
	self.post = self.post.Method("POST",f);
	return self;
}

func (self *Router)setPUT(f func(*Context,string)) *Router{
	self.put = self.put.Method("PUT",f);
	return self;
}

func (self *Router)setDEL(f func(*Context,string)) *Router{
	self.del = self.del.Method("DELETE",f);
	return self;
}

func (self *Router)Do(ctx *Context) int{
	if(!self.get.IsEmpty()){
		if(self.get.Do(ctx) == 0){
			return 0;
		}
	}
	if(!self.post.IsEmpty()){
		if(self.post.Do(ctx) == 0){
			return 0;
		}
	}
	if(!self.put.IsEmpty()){
		if(self.put.Do(ctx) == 0){
			return 0;
		}
	}
	if(!self.del.IsEmpty()){
		if(self.del.Do(ctx) == 0){
			return 0;
		}
	}
	return 1;
}