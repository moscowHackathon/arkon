package controllers

import (
	"github.com/astaxie/beego"
)

type AnswerResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`
}

type AnswerController struct {
	beego.Controller
}

func (c *AnswerController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = AnswerResponse{ID: id, Message: "question"}
	c.ServeJSON()
}
