package routers

import (
	"github.com/moskowhackaton/arkon/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
