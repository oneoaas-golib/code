package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["Content"] = "page not found"
	this.TplNames = "error/404.html"
}

func (this *ErrorController) ErrorDb() {
	this.Data["Content"] = "database is down"
	this.TplNames = "error/db.html"
}

/* End of file  : ErrorController.go */
/* Location     : ./controllers/ErrorController.go */
