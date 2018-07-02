package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type XRoleUser struct { //用户信息表
	Id          uint   `orm:"pk";"auto"` //用户表主键
	UserNo      string `orm:"size(32)"`  //用户表id
	Account     string `orm:"size(32)"`
	Password    string `orm:"size(32)"`
	User_name   string `orm:"size(32)"`
	Postion_no  string `orm:"size(32)"`
	User_sex    uint8
	User_phone  string `orm:"size(11)"`
	User_email  string `orm:"size(32)"`
	User_qq     string `orm:"size(10)"`
	User_idcard string `orm:"size(18)"`
	Status      uint8
	Ctime       int32
	Status_user string `orm:"size(32)"`
	Creator_no  string `orm:"size(32)"`
}

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@"+mysqlurls+"/"+mysqldb+"?charset=utf8", 30, 30)
	beego.Info("链接数据库成功")

	orm.RegisterModel(new(XRoleUser))
	orm.RunSyncdb("default", false, false)
}
