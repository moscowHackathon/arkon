package controllers

import (
	"github.com/astaxie/beego"
)

type StartResponse struct {
	ID string
}

type StartController struct {
	beego.Controller
}

func (c *StartController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = StartResponse{id}
	c.ServeJSON()
}
