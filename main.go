package main

import (
	"github.com/astaxie/beego"
	_ "github.com/moskowHackathon/arkon/routers"
	"github.com/moskowHackathon/arkon/model"
)

func main() {
	beego.AddAPPStartHook(model.InitCore)
	beego.Run()
}
