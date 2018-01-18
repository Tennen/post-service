package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Post struct{
	ClientId string `bson:"clientId" json:"clientId"`
	Topic string `bson:"topic" json:"topic"`
	TopicType string `bson:"topicType" json:"topicType"`
	Content string `bson:"content" json:"content"`
	Author string `bson:"author" json:"author"`
	Receiver string `bson:"receiver" json:"receiver"`
	CreatedAt int64 `bson:"createdAt" json:"createdAt"`
	ReadAt int64 `bson:"readAt" json:"readAt"`
}

type Posts []Post

var c *mgo.Collection

func init() {
	dburi := beego.AppConfig.String("DB_URI")
	session, err := mgo.Dial(dburi)
	if err != nil {
		log.Println(err)
	}
	c = session.DB("post").C("col")
}

func AddOne(post Post) (error){
	data := bson.M{
			"clientId": post.ClientId,
			"topic": post.Topic,
			"topicType": post.TopicType,
			"content": post.Content,
			"author": post.Author,
			"receiver": post.Receiver,
			"createdAt": post.CreatedAt,
			"readAt": post.ReadAt,
		}
	err := c.Insert(data)
	return err
}

func GetAll(topic string) (Posts){
	res := Posts{}
	err := c.Find(bson.M{"topic": topic}).All(&res)
	if err != nil {
		panic(err)
	}
	return res
}
