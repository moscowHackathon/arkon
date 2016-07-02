package controllers

import (
	"github.com/astaxie/beego"
)

type StartResponse struct {
	ID int
}

type StartController struct {
	beego.Controller
}

func (c *StartController) Get() {
	c.Data["json"] = StartResponse{7}
	c.ServeJSON()
}
