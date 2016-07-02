package controllers

import (
	"github.com/astaxie/beego"
)

type CompleteResponse struct {
	ID string
}

type CompleteController struct {
	beego.Controller
}

func (c *CompleteController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = CompleteResponse{id}
	c.ServeJSON()
}
