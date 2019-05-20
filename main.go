package main

import (
	_ "DailyFresh/models"
	_ "DailyFresh/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	//注册模板函数
	beego.AddFuncMap("ShowPrePage", ShowPrePage)
	beego.AddFuncMap("ShowNextPage", ShowNextPage)
	beego.Run()
}

//后台定义一个函数
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func ShowNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}
