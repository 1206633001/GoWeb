package routers

import (
	"github.com/astaxie/beego"
	"ihome/controllers"
)

func init() {
	//APP应用层API路由
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/user/*.go", &controllers.Admin1Controller{})
}
