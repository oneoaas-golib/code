package main

import (
	"code/models"
	_ "code/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	beego.SessionOn = true
	beego.Run()
}
