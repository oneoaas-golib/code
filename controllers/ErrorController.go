package controllers

import (
	"github.com/astaxie/beego"
)

//错误处理控制器
type ErrorController struct {
	beego.Controller
}

//显示404页面
func (this *ErrorController) Error404() {
	this.Data["Content"] = "page not found"
	this.TplNames = "error/404.html"
}

//显示数据库错误页面
func (this *ErrorController) ErrorDb() {
	this.Data["Content"] = "database is down"
	this.TplNames = "error/db.html"
}

/* End of file  : ErrorController.go */
/* Location     : ./controllers/ErrorController.go */
