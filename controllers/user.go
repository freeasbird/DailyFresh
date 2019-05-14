package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowReg() {
	this.TplName = "home/user/register.html"
}

func (this *UserController) HandleReg() {

}
