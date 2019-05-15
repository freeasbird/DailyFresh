package controllers

import (
	"DailyFresh/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"regexp"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowReg() {
	this.TplName = "home/user/register.html"
}

func (this *UserController) HandleReg() {
	//1.获取数据
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")
	//2.检验数据
	if username == "" || pwd == "" || cpwd == "" || email == "" {
		this.Data["errmssg"] = "填写数据不完整,请重新注册"
		this.TplName = "home/user/register.html"
		return
	}
	if pwd != cpwd {
		this.Data["errmssg"] = "两次输入密码不一致,请重新注册"
		this.TplName = "home/user/register.html"
		return
	}
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`)
	res := reg.FindString(email)
	if res == "" {
		this.Data["errmssg"] = "邮箱格式不正确,请重新注册"
		this.TplName = "home/user/register.html"
		return
	}

	//3.处理数据
	o := orm.NewOrm()
	var user models.User
	user.Name = username
	user.Password = pwd
	user.Email = email
	_, err := o.Insert(&user)
	if err != nil {
		this.Data["errmssg"] = "用户名已存在,请重新注册"
		this.TplName = "home/user/register.html"
		return
	}
	//发送邮件
	//配置
	emailConfig := `{"username":"1138894663@qq.com","password":"zsd13715277993","host":"smtp.163.com","port":25}`
	emailConn := utils.NewEMail(emailConfig)
	//发件人
	emailConn.From = "天天生鲜项目系统注册服务"
	//收件人邮箱
	emailConn.To = []string{email}
	//邮件标题
	emailConn.Subject = "天天生鲜用户注册"
	//发送给用户激活地址
	emailConn.Text = "127.0.0.1:8080/active?id=" + strconv.Itoa(user.Id)
	//发送
	emailConn.Send()
	//4.返回视图
	this.Ctx.WriteString("注册成功，请去相应邮箱激活用户！")

}
