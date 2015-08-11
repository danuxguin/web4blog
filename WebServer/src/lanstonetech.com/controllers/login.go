package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"lanstonetech.com/common"
	"lanstonetech.com/models"
)

type LoginController struct {
	beego.Controller
}

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go globalSessions.GC()
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("account", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.DelSession("account")
		//c.Redirect("/", 301)
	}
	c.TplNames = "login.html"
}

func (c *LoginController) Post() {
	account := c.Input().Get("account")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"

	//数据库比较
	pass, err := models.VerifyUser(common.MakeMD5(account), common.MakeMD5(password))
	if err != nil {
		c.Redirect("/login", 301)
		beego.Error(err)
		return
	}
	if pass {
		beego.Error("login successful!")
		maxage := 0
		if autologin {
			maxage = 1<<31 - 1
		}
		c.Ctx.SetCookie("account", common.MakeMD5(account), maxage, "/")
		c.Ctx.SetCookie("password", common.MakeMD5(password), maxage, "/")

		c.SetSession("account", common.MakeMD5(account))
		// c.SetSession("password", password)
	} else {
		c.Redirect("/login", 301)
		beego.Error("login failed!")
		return
	}
	c.Redirect("/", 301)
	return
}

// func verify_user(account, password string) (bool, error) {
// 	var users []models.StoneUser
// 	_, err := models.ORM.Raw("SELECT * from stone_user").QueryRows(&users)
// 	if err != nil {
// 		beego.Error(err)
// 		return false, nil
// 	}
//
// 	for _, user := range users {
// 		if (user.Email == account || user.Telphone == account) && user.Password == password {
// 			return true, nil
// 		}
// 	}
//
// 	return false, nil
// }

// func checkAccount(ctx *context.Context) bool {
// 	// ck, err := ctx.Request.Cookie("account")
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// 	return false
// 	// }
// 	// account := ck.Value
// 	//
// 	// ck, err = ctx.Request.Cookie("password")
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// 	return false
// 	// }
// 	// password := ck.Value
// 	//
// 	// pass, err := models.VerifyUser(account, password)
// 	// if err != nil {
// 	// 	return false
// 	// }
//
// 	return pass
// }
