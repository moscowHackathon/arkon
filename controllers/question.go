package controllers

import (
	"github.com/astaxie/beego"
)

type QuestionResponse struct {
	ID string
}

type QuestionController struct {
	beego.Controller
}

func (c *QuestionController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = QuestionResponse{id}
	c.ServeJSON()
}
