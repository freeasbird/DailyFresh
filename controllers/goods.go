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
	this.TplName = "home/goods/index.html"
}
