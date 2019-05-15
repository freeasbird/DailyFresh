package routers

import (
	"DailyFresh/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, filterFunc)
	//用户注册
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg")
	//用户激活
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
	//用户登录
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//首页
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")

}

var filterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
		return
	}
}
