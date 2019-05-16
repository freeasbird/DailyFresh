package controllers

import (
	"DailyFresh/helper"
	"DailyFresh/models"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"regexp"
	"strconv"
)

type UserController struct {
	beego.Controller
}

//展示注册页面
func (this *UserController) ShowReg() {
	this.TplName = "home/user/register.html"
}

//处理注册请求
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

	var user models.User
	//检验邮箱是否被注册
	o := orm.NewOrm()
	result := o.QueryTable("user").Filter("email", email).One(&user)
	if result != orm.ErrNoRows {
		this.Data["errmssg"] = "该邮箱已被注册"
		this.TplName = "home/user/register.html"
		return
	}

	//3.处理数据
	user.Email = email
	user.Name = username
	user.Password = helper.GetMD5Encode(pwd)
	_, err := o.Insert(&user)
	if err != nil {
		this.Data["errmssg"] = "用户名已存在,请重新注册"
		this.TplName = "home/user/register.html"
		return
	}

	//发送邮件
	//配置smtp服务器账号密码
	emailConfig := `{"username":"1138894663@qq.com","password":"授权码","host":"smtp.qq.com","port":587}`
	emailConn := utils.NewEMail(emailConfig)
	//发件人
	emailConn.From = "1138894663@qq.com"
	//收件人邮箱
	emailConn.To = []string{email}
	//邮件标题
	emailConn.Subject = "天天生鲜用户注册"
	//发送给用户激活地址
	emailConn.Text = "127.0.0.1:8080/active?id=" + strconv.Itoa(user.Id)
	//发送
	send := emailConn.Send()
	if send != nil {
		fmt.Println("邮件发送失败: ", send)
	}
	//4.返回视图
	this.Ctx.WriteString("注册成功，请去相应邮箱激活用户！")
}

//处理用户激活
func (this *UserController) ActiveUser() {
	//1.获取数据
	id, err := this.GetInt("id")
	//2.检验数据
	if err != nil {
		this.Data["errmsg"] = "要激活的用户不存在"
		this.TplName = "home/user/register.html"
		return
	}
	//3.处理数据
	//3.1查询用户操作
	o := orm.NewOrm()
	var user models.User
	user.Id = id
	err = o.Read(&user)
	if err != nil {
		this.Data["errmsg"] = "要激活的用户不存在"
		this.TplName = "home/user/register.html"
		return
	}
	//3.2更新用户操作
	user.Active = true
	o.Update(&user)
	//4.返回视图
	this.Redirect("/login", 302)
}

//展示用户登录页面
func (this *UserController) ShowLogin() {
	userName := this.Ctx.GetCookie("userName")
	//base64解密
	temp, _ := base64.StdEncoding.DecodeString(userName)

	if string(temp) == "" {
		this.Data["userName"] = ""
		this.Data["checked"] = ""
		this.TplName = "home/user/login.html"

	} else {
		this.Data["userName"] = string(temp)
		this.Data["checked"] = "checked"
	}
	this.TplName = "home/user/login.html"
}

//处理用户登录请求
func (this *UserController) HandleLogin() {
	//1.获取数据
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	//2.检验数据
	if username == "" || pwd == "" {
		this.TplName = "home/user/login.html"
		this.Data["errmsg"] = "输入数据不能为空"
		return
	}
	//3.处理数据
	o := orm.NewOrm()
	var user models.User
	result := o.QueryTable("user").Filter("name", username).Filter("password", helper.GetMD5Encode(pwd)).One(&user)
	if result == orm.ErrNoRows {

		this.TplName = "home/user/login.html"
		this.Data["errmsg"] = "账号或密码错误"
		return
	}

	if user.Active == false {
		this.TplName = "home/user/login.html"
		this.Data["errmsg"] = "该账户未激活,请去邮箱激活"
		return
	}

	//4.返回视图 跳转主页
	remember := this.GetString("remember")
	if remember == "on" {
		//base64加密
		temp := base64.StdEncoding.EncodeToString([]byte(username))
		this.Ctx.SetCookie("userName", temp, 24*3600)
	} else {
		this.Ctx.SetCookie("userName", "", -1)
	}
	this.SetSession("userName", username)
	this.Redirect("/", 302)

}

//退出登陆
func (this *UserController) Logout() {
	this.DelSession("userName")
	//跳转登录页面
	this.Redirect("/login", 302)
}

//展示用户中心个人信息页面
func (this *UserController) ShowUserCenterInfo() {
	username := GetUser(&this.Controller)
	this.Data["userName"] = username

	var addr models.Address
	o := orm.NewOrm()
	o.QueryTable("Address").RelatedSel("User").Filter("user__name", username).Filter("Is_default", true).One(&addr)
	if addr.Id == 0 {
		this.Data["addr"] = ""
	} else {
		this.Data["addr"] = addr
	}
	this.Layout = "home/user/userCenterLayout.html"
	this.TplName = "home/user/user_center_info.html"
}

//展示用户中心订单信息页面
func (this *UserController) ShowUserCenterOrder() {
	GetUser(&this.Controller)
	this.Layout = "home/user/userCenterLayout.html"
	this.TplName = "home/user/user_center_order.html"
}

//展示用户中地址页面
func (this *UserController) ShowUserCenterSite() {
	userName := GetUser(&this.Controller)
	//获取地址信息
	o := orm.NewOrm()
	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Name", userName).Filter("Is_default", true).One(&addr)
	//传递数据
	this.Data["addr"] = addr
	this.Layout = "home/user/userCenterLayout.html"
	this.TplName = "home/user/user_center_site.html"
}

//处理用户中心地址数据
func (this *UserController) HandleUserCenterSite() {
	//1.获取数据
	receiver := this.GetString("receiver")
	addr := this.GetString("addr")
	zipCode := this.GetString("zipCode")
	phone := this.GetString("phone")
	//2.校验数据
	if receiver == "" || addr == "" || zipCode == "" || phone == "" {
		beego.Info("添加数据不完整")
		this.Redirect("/user/userCenterSite", 302)
		return
	}
	//3.处理数据
	//插入地址数据
	o := orm.NewOrm()
	var addrUser models.Address
	addrUser.Is_default = true
	err := o.Read(&addrUser, "Is_default")
	//把旧默认地址改为非默认地址
	if err == nil {
		addrUser.Is_default = false
		o.Update(&addrUser)

	}
	//插入新地址并作为默认地址
	//关联
	userName := this.GetSession("userName")
	var user models.User
	user.Name = userName.(string)
	o.Read(&user, "Name")

	var addrUserNew models.Address
	addrUserNew.Receiver = receiver
	addrUserNew.Addr = addr
	addrUserNew.Zip_code = zipCode
	addrUserNew.Phone = phone
	addrUserNew.Is_default = true
	addrUserNew.User = &user
	o.Insert(&addrUserNew)
	//4.返回视图
	this.Redirect("/user/userCenterSite", 302)
}
