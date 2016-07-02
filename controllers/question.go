package controllers

import (
	"github.com/astaxie/beego"
	"github.com/moskowHackathon/arkon/models"
)

type QuestionResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type QuestionController struct {
	beego.Controller
}

func (c *QuestionController) Get() {
	id := c.Ctx.Input.Param(":id")

	s, err := models.GetSession(id)

	if err != nil {
		c.Data["json"] = QuestionResponse{
			ID:    id,
			Error: err.Error(),
		}
		c.ServeJSON()
		return
	}

	q := s.GetQuestion()
	c.Data["json"] = QuestionResponse{
		ID:      id,
		Message: q,
	}
	c.ServeJSON()
	return
}
