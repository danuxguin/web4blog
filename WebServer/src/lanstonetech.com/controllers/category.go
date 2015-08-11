package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lanstonetech.com/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")

	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301)
		return
	}

	// account := fmt.Sprintf("%v", c.GetSession("account"))
	// password := fmt.Sprintf("%v", c.GetSession("password"))
	// pass, err := models.VerifyUser(account, password)
	// if err != nil {
	// 	beego.Error(err)
	// 	return
	// }

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)
	if len(name) == 0 {
		name = "游客"
	}
	c.Data["Name"] = name

	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsLogin"] = pass
	c.TplNames = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
