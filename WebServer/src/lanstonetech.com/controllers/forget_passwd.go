package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lanstonetech.com/models"
	"os/exec"
)

type ForgetPasswdController struct {
	beego.Controller
}

func (c *ForgetPasswdController) Get() {
	c.TplNames = "forget_passwd.html"
}

func (c *ForgetPasswdController) Post() {
	telphone := c.Input().Get("telphone")
	email := c.Input().Get("email")

	new_password := GetRandomSring(6)
	succ, err := models.UpdateUserPassword(telphone, email, new_password)
	if err != nil {
		beego.Error(err)
		c.Ctx.WriteString("验证错误")
		return
	}

	if succ {
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo \"你的初始化密码为:%s\"| mutt -s \"密码重置-蓝石头\" %s", new_password, email))
		_, err := cmd.Output()
		if err != nil {
			beego.Error(err)
			c.Ctx.WriteString("验证错误")
		}

		c.Ctx.WriteString("密码已重置，请查看注册邮箱")
		return
	}

	c.Ctx.WriteString("非注册用户")
}
