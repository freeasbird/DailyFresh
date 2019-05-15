package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

//展示首页
func (this *IndexController) ShowIndex() {
	this.TplName = "home/index/index.html"
}
