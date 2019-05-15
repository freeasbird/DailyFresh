package controllers

import (
	"github.com/astaxie/beego"
)

//商品控制器
type GoodsController struct {
	beego.Controller
}

//展示商品首页
func (this *GoodsController) ShowIndex() {
	userName := this.GetSession("userName")
	if userName == nil {
		this.Data["userName"] = ""
	} else {
		this.Data["userName"] = userName.(string)

	}
	this.TplName = "home/goods/index.html"
}
