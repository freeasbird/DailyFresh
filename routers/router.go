package routers

import (
	"DailyFresh/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg")

}
