package controllers

import (
	"DailyFresh/models"
	"bytes"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"math"
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

	adminName := GetAdminName(&this.Controller)
	typeName := this.GetString("select")

	//每页记录数
	pageSize := 2

	//获取页码
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}
	//获取起始数据位置
	start := (pageIndex - 1) * pageSize

	o := orm.NewOrm()
	qs := o.QueryTable("GoodsSku")
	var count int64
	var perSize int64 = 10
	if typeName == "" {
		count, _ = qs.Count()

	} else {
		count, _ = qs.Limit(perSize, start).RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).Count()
	}
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	var goods []models.GoodsSku
	var types []models.GoodsType
	//获取数据
	conn, err := redis.Dial("tcp", ":6379")
	//从redis中获取数据
	//解码
	rep, err := conn.Do("get", "types")
	data, err := redis.Bytes(rep, err)
	//获取解码器
	dec := gob.NewDecoder(bytes.NewReader(data))
	dec.Decode(&types)
	if len(types) == 0 {
		//从redis中获取数据不成功,从mysql获取数据
		o.QueryTable("GoodsType").All(&types)
		//把获取到的数据存储到redis中
		//编码操作
		var buffer bytes.Buffer
		//获取编码器
		enc := gob.NewEncoder(&buffer)
		//编码
		enc.Encode(&types)
		//存入redis
		conn.Do("SET", "types", buffer.Bytes())
		beego.Info("从mysql中获取数据")
	}
	//传递数据
	this.Data["types"] = types
	this.Data["userName"] = adminName
	this.Data["typeName"] = typeName
	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["goodslist"] = goods

	//指定模板
	this.TplName = "admin/goods/goodsList.html"

}

//展示商品类型添加页面
func (this *GoodsController) ShowAdminGoodsTypeAdd() {
	GetAdminName(&this.Controller)
	o := orm.NewOrm()
	var types []models.GoodsType
	o.QueryTable("GoodsType").All(&types)

	//传递数据
	this.Data["types"] = types
	this.TplName = "admin/goods/goodsTypeAdd.html"
}

//处理商品类型添加
func (this *GoodsController) HandleAdminGoodsTypeAdd() {

}
