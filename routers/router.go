package routers

import (
	"DailyFresh/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//前台模块
	ns1 := beego.NewNamespace("/home",
		beego.NSNamespace("/user",
			//路由拦截
			beego.NSBefore(filterHomeFunc),
			//用户注册
			beego.NSRouter("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg"),
			//用户登录
			beego.NSRouter("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin"),
			//用户激活
			beego.NSRouter("/active", &controllers.UserController{}, "get:ActiveUser"),
			//退出登陆
			beego.NSRouter("/logout", &controllers.UserController{}, "get:Logout"),
			//用户中心个人信息页
			beego.NSRouter("/userCenterInfo", &controllers.UserController{}, "get:ShowUserCenterInfo"),
			//用户中心订单信息页
			beego.NSRouter("/userCenterOrder", &controllers.UserController{}, "get:ShowUserCenterOrder"),
			//用户中心地址页
			beego.NSRouter("/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite;post:HandleUserCenterSite"),
		),
	)

	//后台模块
	ns2 := beego.NewNamespace("/admin",
		beego.NSNamespace("/user",
			beego.NSBefore(filterAdminFunc),
			//后台注册页
			beego.NSRouter("/register", &controllers.UserController{}, "get:ShowAdminReg;post:HandleAdminReg"),
			//后台登陆页
			beego.NSRouter("/login", &controllers.UserController{}, "get:ShowAdminLogin;post:HandleAdminLogin"),
			//后台主页
			beego.NSRouter("/index", &controllers.UserController{}, "get:ShowAdminIndex"),
		),
		beego.NSNamespace("/goods",
			beego.NSBefore(filterAdminFunc),
			beego.NSRouter("/goodsList", &controllers.GoodsController{}, "get:ShowAdminGoodsList"),
			beego.NSRouter("/goodsType", &controllers.GoodsController{}, "get:ShowAdminGoodsType"),
			beego.NSRouter("/goodsTypeAdd", &controllers.GoodsController{}, "get:ShowAdminGoodsTypeAdd;post:HandleAdminGoodsTypeAdd"),
			beego.NSRouter("/goodsTypeDel", &controllers.GoodsController{}, "get:HandleAdminGoodsTypeDel"),
			beego.NSRouter("/goodsTypeEdit", &controllers.GoodsController{}, "get:ShowAdminGoodsTypeEdit;post:HandleAdminGoodsTypeEdit"),
		),
	)

	//首页
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")

	//注册路由
	beego.AddNamespace(ns1, ns2)
}

//前台模块路由拦截函数
var filterHomeFunc = func(ctx *context.Context) {
	path := ctx.Request.URL.Path
	//无需拦截的路由使用map查询效率高
	allowPathMap := make(map[string]int)
	allowPathMap["/home/user/login"] = 1
	allowPathMap["/home/user/register"] = 1
	if allowPathMap[path] == 1 {
		return
	} else {
		userName := ctx.Input.Session("userName")
		if userName == nil {
			ctx.Redirect(302, "/home/user/login")
			return
		}
	}
}

//后台模块路由拦截函数
var filterAdminFunc = func(ctx *context.Context) {
	path := ctx.Request.URL.Path
	allowPathMap := make(map[string]int)
	allowPathMap["/admin/user/login"] = 1
	allowPathMap["/admin/user/register"] = 1
	if allowPathMap[path] == 1 {
		return
	} else {
		userName := ctx.Input.Session("adminName")
		fmt.Println(userName)
		if userName == nil {
			ctx.Redirect(302, "/admin/user/login")
			return
		}
	}
}
