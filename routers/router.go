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
	//退出登陆
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	//用户中心个人信息页
	beego.Router("/user/userCenterInfo", &controllers.UserController{}, "get:ShowUserCenterInfo")
	//用户中心订单信息页
	beego.Router("/user/userCenterOrder", &controllers.UserController{}, "get:ShowUserCenterOrder")
	//用户中心地址页
	beego.Router("/user/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite;post:HandleUserCenterSite")

}

var filterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
		return
	}
}
