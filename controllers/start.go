package controllers

import (
	"github.com/astaxie/beego"
	"github.com/moskowHackathon/arkon/models"
)

type StartResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type StartController struct {
	beego.Controller
}

func (c *StartController) Get() {
	id := c.Ctx.Input.Param(":id")

	models.NewSession(id)

	c.Data["json"] = StartResponse{ID: id, Message: "ok"}
	c.ServeJSON()
	return
}
