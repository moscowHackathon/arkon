package controllers

import (
	"github.com/astaxie/beego"
)

type CompleteResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`
}

type CompleteController struct {
	beego.Controller
}

func (c *CompleteController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = CompleteResponse{ID: id, Message: "ok"}
	c.ServeJSON()
}
