package controllers

import (
	"github.com/astaxie/beego"
)

type StartResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`
}

type StartController struct {
	beego.Controller
}

func (c *StartController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = StartResponse{ID: id, Message: "ok"}
	c.ServeJSON()
}
