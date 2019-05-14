package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//用户表
type User struct {
	Id        int
	Name      string       `orm:"size(20);description(用户名)"`                 //用户名
	Password  string       `orm:"size(50);description(登陆密码)"`                //登陆密码
	Email     string       `orm:"size(50);description(邮箱)"`                  //邮箱
	Active    bool         `orm:"default(false);description(是否激活)"`          //是否激活
	Power     int          `orm:"default(0);description(权限设置 0表示未激活 1表示激活)"` //权限设置 0表示未激活 1表示激活
	Address   []*Address   `orm:"reverse(many)"`
	OrderInfo []*OrderInfo `orm:"reverse(many)"`
}

//地址表
type Address struct {
	Id         int
	Receiver   string       `orm:"size(20);description(收件人)"`                   //收件人
	Addr       string       `orm:"size(50);description(收件地址)"`                  //收件地址
	Zip_code   string       `orm:"size(20);description(邮编)"`                    //邮编
	Phone      string       `orm:"size(20);description(联系电话)"`                  //联系电话
	Is_default bool         `orm:"default(false);description(是否默认 0为非默认 1为默认)"` //是否默认 0为非默认 1为默认
	User       *User        `orm:"rel(fk);description(用户ID)"`                   //用户ID
	OrderInfo  []*OrderInfo `orm:"reverse(many)"`
}

//商品类型表
type GoodsType struct {
	Id                   int
	Name                 string                  `orm:"description(类型名称)"` //类型名称
	Logo                 string                  `orm:"description(logo)"` //logo
	Image                string                  `orm:"description(图片)"`   //图片
	GoodsSku             []*GoodsSku             `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

//商品SPU表
type Goods struct {
	Id       int
	Name     string      `orm:"size(20);description(商品名称)"`    //商品名称
	Detail   string      `orm:"size(200);description(商品详情描述)"` //商品详情描述
	GoodsSku []*GoodsSku `orm:"reverse(many)"`
}

//商品SKU表
type GoodsSku struct {
	Id                   int
	Goods                *Goods                  `orm:"rel(fk);description(商品SPU)"`                          //商品SPU
	GoodsType            *GoodsType              `orm:"rel(fk);description(商品所属类型)"`                         //商品所属类型
	Name                 string                  `orm:"description(商品名称)"`                                   //商品名称
	Desc                 string                  `orm:"description(商品简介)"`                                   //商品简介
	Price                float64                 `orm:"digits(12);decimals(2);default(0);description(商品价格)"` //商品价格
	Unite                string                  `orm:"description(商品单位)"`                                   //商品单位
	Image                string                  `orm:"description(商品图片)"`                                   //商品图片
	Stock                uint                    `orm:"default(0);description(商品库存)"`                        //商品库存
	Sales                int                     `orm:"default(0);description(商品销量)"`                        //商品销量
	Status               int                     `orm:"default(1);description(商品状态)"`                        //商品状态
	Time                 time.Time               `orm:"auto_now_add;description(添加时间)"`                      //添加时间
	GoodsImage           []*GoodsImage           `orm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
	OrderGoods           []*OrderGoods           `orm:"reverse(many)"`
}

//商品图片表
type GoodsImage struct {
	Id       int
	Image    string    `orm:"description(商品图片)"` //商品图片
	GoodsSku *GoodsSku `orm:"rel(fk)"`           //商品SKU
}

//首页轮播商品展示表
type IndexGoodsBanner struct {
	Id       int
	GoodsSku *GoodsSku `orm:"rel(fk);description(商品SKU)"`   //商品SKU
	Image    string    `orm:"description(商品图片)"`            //商品图片
	Index    int       `orm:"default(0);description(展示顺序)"` //展示顺序
}

//首页分类商品展示表
type IndexTypeGoodsBanner struct {
	Id           int
	GoodsType    *GoodsType `orm:"rel(fk);description(商品类型)"`                //商品类型
	GoodsSku     *GoodsSku  `orm:"rel(fk);description(商品SKU)"`               //商品SKU
	Display_Type int        `orm:"default(1);description(展示类型 0代表标题 1代表文字)"` //展示类型 0代表标题 1代表文字
	Index        int        `orm:"default(0);description(展示顺序)"`             //展示顺序
}

//首页促销商品展示表
type IndexPromotionBanner struct {
	Id    int
	Name  string `orm:"size(20);description(活动名称)"`   //活动名称
	Url   string `orm:"size(50);description(活动链接)"`   //活动链接
	Image string `orm:"description(商品所属类型)"`          //活动图片
	Index int    `orm:"default(0);description(展示顺序)"` //展示顺序
}

//订单表
type OrderInfo struct {
	Id            int
	OrderId       string        `orm:"unique;description(订单号)"`                  //订单号
	User          *User         `orm:"rel(fk);description(用户)"`                  //用户
	Address       *Address      `orm:"rel(fk);description(地址)"`                  //地址
	Pay_Method    int           `orm:"description(付款方式)"`                        //付款方式
	Total_Count   int           `orm:"default(1);description(商品数量)"`             //商品数量
	Total_Price   float64       `orm:"digits(12);decimals(2);description(商品总价)"` //商品总价
	Transit_Price float64       `orm:"digits(12);decimals(2);description(运费)"`   //运费
	Order_status  int           `orm:"default(1);description(订单状态)"`             //订单状态
	Trade_No      string        `orm:"description(支付编号)"`                        //支付编号
	Time          time.Time     `orm:"auto_now_add;description(添加时间)"`           //添加时间
	OrderGoods    []*OrderGoods `orm:"reverse(many)"`
}

//OrderInfo表引擎设置为INNODB
func (o *OrderInfo) TableEngine() string {
	return "INNODB"
}

type OrderGoods struct {
	Id        int
	OrderInfo *OrderInfo `orm:"rel(fk);description(订单)"`                  //订单
	GoodsSku  *GoodsSku  `orm:"rel(fk);description(商品)"`                  //商品
	Count     int        `orm:"default(1);description(商品数量)"`             //商品数量
	Price     float64    `orm:"digits(12);decimals(2);description(商品价格)"` //商品价格
	Comment   string     `orm:"description(评论)"`                          //评论
}

func init() {
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/daliy_fresh?charset=utf8")
	//注册模型
	orm.RegisterModel(new(User), new(Address), new(GoodsType), new(Goods), new(GoodsSku), new(GoodsImage), new(IndexGoodsBanner), new(IndexTypeGoodsBanner), new(IndexPromotionBanner), new(OrderInfo), new(OrderGoods))
	//创建表
	orm.RunSyncdb("default", false, true)

}
