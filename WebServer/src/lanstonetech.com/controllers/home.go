package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lanstonetech.com/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplNames = "home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsLogin"] = pass

	topics, err := models.GetAllTopics(true, "")
	if err != nil {
		beego.Error(err)
		return
	}

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)
	if len(name) == 0 {
		name = "游客"
	}

	c.Data["Name"] = name
	c.Data["Topics"] = topics
}
