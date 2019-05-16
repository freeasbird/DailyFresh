package controllers

import (
	"github.com/astaxie/beego"
)

//商品控制器
type GoodsController struct {
	beego.Controller
}

//获取用户
func GetUser(this *beego.Controller) string {
	userName := this.GetSession("userName")
	if userName == nil {
		this.Data["userName"] = ""
		return ""
	} else {
		this.Data["userName"] = userName.(string)
		return userName.(string)
	}
}

//展示商品首页
func (this *GoodsController) ShowIndex() {
	GetUser(&this.Controller)
	this.TplName = "home/goods/index.html"
}
