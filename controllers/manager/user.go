package manager

import (
	"code/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	var err error
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/user_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/user_index_heade.html"
	this.Data["Users"], err = models.GetUsers("1", 10)
	if err != nil {
		beego.Error(err)
	}
}

func (this *UserController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		this.Layout = "manager/lauout.html"
		this.TplNames = "manager/user_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/user_create_heade.html"
		return
	}

	name := this.GetString("name")
	passone := this.GetString("passone")
	passtwo := this.GetString("passtwo")
	if name == "" || passone == "" || passtwo == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	if passone != passtwo {
		this.Data["json"] = map[string]string{"code": "error", "info": "两次填写的密码不一致！"}
		this.ServeJson()
		return
	}
}
