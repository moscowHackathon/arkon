package main

import (
	"github.com/astaxie/beego"
	_ "github.com/moskowhackathon/arkon/routers"
)

func main() {
	beego.Run()
}
