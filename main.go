package main

import (
	_ "DailyFresh/models"
	_ "DailyFresh/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true

	beego.Run()
}
