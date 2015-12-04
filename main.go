package main

import (
	"code/controllers"
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
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("/upload", "upload")
	// beego.SetLogger("file", `{"filename":"./logs/logs.log"}`)
	// beego.BeeLogger.DelLogger("console")
	//注册模板函数
	beego.AddFuncMap("getarticletrash", GetArticleTrashCount)
	beego.Run()
}

func GetArticleTrashCount() int64 {
	count, _ := models.GetArticleCount([]int{0})
	return count
}
