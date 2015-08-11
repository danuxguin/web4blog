package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"lanstonetech.com/models"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	beego.Error("hello signup")
	c.TplNames = "signup.html"
}

func (c *SignupController) Post() {
	var user models.StoneUser
	user.Name = c.Input().Get("username")
	user.Email = c.Input().Get("email")
	user.Telphone = c.Input().Get("telphone")
	user.Password = c.Input().Get("password")

	err := models.AddUser(user)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/login", 301)
}
