package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Data["Website"] = "admin"
	this.Data["Email"] = "test1232"
	this.TplName = "index.tpl"
}

type Admin1Controller struct {
	beego.Controller
}

func (this *Admin1Controller) Get() {
	this.Data["Website"] = "admin1111"
	this.Data["Email"] = "test1232111"
	this.Ctx.WriteString(this.Ctx.Input.Param(":path") + ".go")
	this.TplName = "index.tpl"
}
