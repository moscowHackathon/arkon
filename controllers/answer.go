package controllers

import (
	"github.com/astaxie/beego"
)

type AnswerResponse struct {
	ID int
}

type AnswerController struct {
	beego.Controller
}

func (c *AnswerController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = AnswerResponse{id}
	c.ServeJSON()
}
