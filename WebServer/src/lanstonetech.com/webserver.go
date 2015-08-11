package main

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	_ "lanstonetech.com/models"
	_ "lanstonetech.com/routers"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
