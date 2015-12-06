package main;

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Question struct{
	Key string				`json:"key"`             
	Title string			`json:"title"`
	Description string		`json:"description"`
	Topic string			`json:"topic"`
}

type Answer struct{
	Key string
	Content string	
}

type Mongodb struct{
	Session *mgo.Session
	Collection *mgo.Collection
}

func createMongodb() *Mongodb{
	mongodb := &Mongodb{};
	mongodb.Session = nil;
	mongodb.Collection = nil;
	return mongodb;
}

func (self *Mongodb) ConnectMgo(){
	session,err := mgo.Dial("localhost:27017");
	if(err != nil){
		panic(err);	
	}
	session.SetMode(mgo.Monotonic, true);
	self.Session = session;
}

func (self *Mongodb) CloseMgo(){
	if(self.Session != nil){
		self.Session.Close()
	}
}

func (self *Mongodb) GetCollection(database string,table string){
	self.Collection = self.Session.DB(database).C(table);
}

func (self *Mongodb) InsertQuestion(question *Question){
	err := self.Collection.Insert(question);
	if(err != nil){
		log.Fatal(err);
	}
}

func (self *Mongodb) FindQuestionByKey(key string) *Question{
	result := &Question{};
	err := self.Collection.Find(bson.
							M{"key": key}).
							One(result);
	if(err != nil){
		log.Fatal(err);
	}
	return result;
}

func (self *Mongodb) InsertAnswer(answer *Answer){
	err := self.Collection.Insert(answer);
	if(err != nil){
		log.Fatal(err);
	}
}

func (self *Mongodb) FindAnswerByKey(key string) *Answer{
	result := &Answer{};
	err := self.Collection.Find(bson.
							M{"key": key}).
							One(result);
	if(err != nil){
		log.Fatal(err);
	}
	return result;
}