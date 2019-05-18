package controllers

import (
	"DailyFresh/helper"
	"DailyFresh/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//商品控制器
type GoodsController struct {
	beego.Controller
}

//************************************【前台模块】*******************************************//
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

//************************************【后台模块】*******************************************//

func (this *GoodsController) ShowAdminGoodsList() {
	GetAdminName(&this.Controller)
	typeName := this.GetString("select")
	page, err := this.GetInt64("page")
	if err != nil {
		page = 1
	}
	var goods []models.GoodsSku

	o := orm.NewOrm()
	qs := o.QueryTable("GoodsSku")
	var count int64
	var perSize int64 = 10
	if typeName == "" {
		count, _ = qs.Count()
		_, err := qs.All(&goods)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		count, _ = qs.Limit(perSize, page-1).RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).Count()
		_, err = qs.Limit(perSize, page-1).RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).All(&goods)
		if err != nil {
			fmt.Println(err)
		}
	}

	pg := helper.PageHelperInit(count, 10, page)
	this.Data["pg"] = pg
	this.Data["goodslist"] = goods
	this.Layout = "admin/layout/adminLayout.html"
	this.TplName = "admin/goods/goodsList.html"

}
