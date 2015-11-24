package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) Error404() {
	e.Data["Content"] = "page not found"
	e.TplNames = "error/404.html"
}

func (e *ErrorController) ErrorDb() {
	e.Data["Content"] = "database is down"
	e.TplNames = "error/db.html"
}
