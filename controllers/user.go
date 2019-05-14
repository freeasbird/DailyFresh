package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowReg() {
	this.TplName = "home/user/register.html"
}

func (this *UserController) HandleReg() {
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")
	if username == "" || pwd == "" || cpwd == "" || email == "" {
		this.Data["errmssg"] = "填写数据不完整"
		this.TplName = "home/user/register.html"
	}
	regexp.Compile("")

}
