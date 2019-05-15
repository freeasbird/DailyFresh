package routers

import (
	"DailyFresh/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//用户注册
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg")
	//用户激活
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
	//用户登录
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//首页
	beego.Router("/index", &controllers.UserController{}, "")

}
