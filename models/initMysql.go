package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户区
type XUser struct { //用户信息表
	Id              uint   `orm:"auto"`           //用户表主键
	Phone           string `orm:"size(11);index"` //手机账号
	Password        string `orm:"size(32)"`       //用户密码
	User_name       string `orm:"size(32)"`       //用户名
	Header_img      string `orm:"size(100)"`      //用户头像
	Authority_group uint   //权限组
	User_sex        uint8  //性别：0=保密；1=男；2=女
	Age             uint8  //年龄
	User_email      string `orm:"size(32)"` //用户邮箱
	User_qq         string `orm:"size(10)"` //用户QQNumber
	User_idcard     string `orm:"size(18)"` //用户身份证ID
	Status          uint8  //是否实名认证：0=没认证；1=认证
	Ctime           uint64 //注册时间时间戳
}                            //用户记录表
type XUserAuthority struct { //用户权限组
	Id            uint   `orm:"auto"`     //权限组ID
	Name          string `orm:"size(32)"` //权限组名称
	ChangeInfo    uint8  //用户资料修改  0=有；1=没  描述禁止修改资料
	News          uint8  //禁止消息系统
	PostDemand    uint8  //禁止发布需求（租）
	PostCommodity uint8  //禁止发布商品房（出租）
	PostTz        uint8  //禁止发布帖子和留言
	Login         uint8  //禁止登陆
	Ctime         uint64 //创建时间戳
	Cuser         uint   //创建用户ID
} //普通用户权限表

//经纪人区
type XAgent struct {
	Id       uint   `orm:"auto"`      //经纪人ID
	Phone    string `orm:"size(11)"`  //经纪人登陆账号
	Password string `orm:"size(32)"`  //经纪人登陆密码
	Name     string `orm:"size(32)"`  //经纪人名称
	Email    string `orm:"size(32)"`  //经纪人邮件
	Qq       string `orm:"size(32)"`  //经纪人QQ
	Wechat   string `orm:"size(32)"`  //经纪人微信
	Zp       string `orm:"size(100)"` //经纪人证件照
	Sex      uint8  //性别
	Age      uint8  //年龄
	Idcard   string `orm:"size(18)"` //经纪人身份证号码
	Ctime    uint64 //创建时间
	Cuser    uint   //创建管理员
} //经纪人表

//管理员区
type XAdmin struct {
	Id uint `orm:"auto"`
	Phone
}

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@"+mysqlurls+"/"+mysqldb+"?charset=utf8", 30, 30)
	beego.Info("链接数据库成功")

	orm.RegisterModel(new(XUser), new((XUserAuthority)))
	orm.RunSyncdb("default", false, false)
}
