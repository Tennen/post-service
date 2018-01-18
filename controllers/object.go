package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"post-service/models"
)

// Operations about object
type PostController struct {
	beego.Controller
}

type PostsResponse struct {
	Data models.Posts `json:"data"`
	WjlCode int `json:"wjlCode"`
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *PostController) Post() {
	var p models.Post
	json.Unmarshal(this.Ctx.Input.RequestBody, &p)
	err := models.AddOne(p)
	if err != nil {
		panic(err)
	}
	res := PostsResponse{Data: nil, WjlCode: 2000}
	this.Data["json"] = &res
	this.ServeJSON()
}

func (this *PostController) Get() {
	topic := this.GetString("topic")
	data := models.GetAll(topic)
	res := PostsResponse{Data: data, WjlCode: 2000}
	this.Data["json"] = &res
	this.ServeJSON()
}