package routers

import (
	"DailyFresh/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg")
}
