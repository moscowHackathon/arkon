package main

import (
	"github.com/astaxie/beego"
	_ "github.com/moscowHackathon/arkon/routers"
	"github.com/moscowHackathon/arkon/models"
)

func main() {
	beego.AddAPPStartHook(models.InitCore)
	beego.Run()
}
