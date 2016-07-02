package routers

import (
	"github.com/astaxie/beego"
	"github.com/moskowhackaton/arkon/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/start", &controllers.StartController{})
}
