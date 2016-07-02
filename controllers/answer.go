package controllers

import (
	"github.com/astaxie/beego"
	"github.com/moskowHackathon/arkon/models"
)

type AnswerResponse struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`
}

type AnswerController struct {
	beego.Controller
}

func (c *AnswerController) Post() {
	id := c.Ctx.Input.Param(":id")

	a, err  := c.GetInt("answer", 0)
	if err != nil {
		c.Data["json"] = AnswerResponse{ID: id, Error: err.Error()}
		c.ServeJSON()
		return
	}

	s, err := models.GetSession(id)
	if err != nil {
		c.Data["json"] = AnswerResponse{ID: id, Error: err.Error()}
		c.ServeJSON()
		return
	}

	r := s.Answer(a)
	if r == true {
		c.Data["json"] = AnswerResponse{ID: id, Message: "complete"}
		c.ServeJSON()
		return

	}

	c.Data["json"] = AnswerResponse{ID: id, Message: "question"}
	c.ServeJSON()
	return
}
