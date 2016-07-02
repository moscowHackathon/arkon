package controllers

import (
	"github.com/astaxie/beego"
)

type QuestionResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`

}

type QuestionController struct {
	beego.Controller
}

func (c *QuestionController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = QuestionResponse{
		ID:id,
		Message: "Question 1",
	}
	c.ServeJSON()
}
