package routers

import (
	"github.com/astaxie/beego"
	"github.com/moskowhackathon/arkon/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/start", &controllers.StartController{})
}
