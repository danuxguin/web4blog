package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"lanstonetech.com/models"
	"math/rand"
	"time"
)

func GetCookieValue(c *beego.Controller, name string) (string, error) {
	value, err := c.Ctx.Request.Cookie(name)
	if err != nil {
		return "", err
	}

	return value.Value, nil
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		beego.Error(err)
		return false
	}
	account := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		beego.Error(err)
		return false
	}
	password := ck.Value

	pass, err := models.VerifyUser(account, password)
	if err != nil {
		return false
	}

	return pass
}

func GetRandomSring(num int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	l := len(s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}

	return buf.String()
}
