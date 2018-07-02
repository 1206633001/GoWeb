package main

import (
	"github.com/astaxie/beego"
	_ "ihome/models"
	_ "ihome/routers"
)

func main() {
	beego.Run()
}
