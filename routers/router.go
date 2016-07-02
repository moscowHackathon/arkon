package routers

import (
	"github.com/astaxie/beego"
	"github.com/moskowHackathon/arkon/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/start", &controllers.StartController{})
}
