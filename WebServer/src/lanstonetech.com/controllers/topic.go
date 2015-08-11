package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lanstonetech.com/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.TplNames = "topic.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)

	topics := make([]*models.StoneTopic, 0)
	var err error
	if len(name) == 0 {
		// c.Redirect("/login?exit=true", 302)
		c.Data["Name"] = "游客"
		topics, err = models.GetAllTopics(false, "")
		if err != nil {
			beego.Error(err)
			c.Redirect("/", 301)
			return
		}
	} else {
		c.Data["Name"] = name
		topics, err = models.GetAllTopics(false, account)
		if err != nil {
			beego.Error(err)
			c.Redirect("/", 301)
			return
		}
	}

	c.Data["Topics"] = topics
}

func (c *TopicController) Post() {
	session_account := fmt.Sprintf("%v", c.GetSession("account"))
	session_password := fmt.Sprintf("%v", c.GetSession("password"))
	pass, err := models.VerifyUser(session_account, session_password)
	if err != nil {
		beego.Error(err)
		return
	}

	if pass {
		c.Redirect("/login", 302)
		return
	}

	//account := c.Input().Get("account")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	tid := c.Input().Get("tid")
	op_type := c.Input().Get("type")

	switch op_type {
	case "1":
		//添加文章
		err := models.AddTopic(session_account, title, content)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/topic", 301)
		break
	case "2":
		//修改文章
		err := models.ModifyTopic(session_account, tid, title, content)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect(fmt.Sprintf("/topic/view/%s", tid), 301)
		break
	case "3":
		name, _ := models.GetUserName(session_account)
		err := models.AddTopicReply(tid, -1, session_account, name, content)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect(fmt.Sprintf("/topic/view/%s", tid), 301)
		break
	}
}

func (c *TopicController) Add() {
	// account, err := c.Ctx.Request.Cookie("account")
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/login", 301)
	// 	return
	// }

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)
	if len(name) == 0 {
		c.Redirect("/login?exit=true", 302)
	}
	c.Data["Name"] = name

	c.TplNames = "topic_add.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsLogin"] = pass
	// c.Data["Account"] = account.Value
}

func (c *TopicController) View() {
	c.TplNames = "topic_view.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topic, err := models.GetTopicByID(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
	}

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)
	if len(name) == 0 {
		name = "游客"
	}
	c.Data["Name"] = name

	fmt.Printf("account = %v\ttopic.account = %v\n", account, topic.Account)
	if account == topic.Account {
		c.Data["Ishost"] = true
	} else {
		c.Data["Ishost"] = false
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")

	replys, err := models.GetTopicReplys(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
	}

	c.Data["Replys"] = replys
}

func (c *TopicController) Modify() {
	// account, err := c.Ctx.Request.Cookie("account")
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Redirect("/login", 301)
	// 	return
	// }

	c.TplNames = "topic_modify.html"
	tid := c.Input().Get("tid")
	//c.Data["IsLogin"] = checkAccount(c.Ctx)
	topic, err := models.GetTopicByID(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	account := fmt.Sprintf("%v", c.GetSession("account"))
	name, _ := models.GetUserName(account)
	if len(name) == 0 {
		c.Redirect("/login?exit=true", 302)
	}
	c.Data["Name"] = name

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
	// c.Data["Account"] = account.Value
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		// if pass {
		c.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/", 302)
	return
}
