package routers

import (
	"github.com/astaxie/beego"
	"lanstonetech.com/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/forgetpasswd", &controllers.ForgetPasswdController{})
}
