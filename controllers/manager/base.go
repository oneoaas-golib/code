package manager

import (
	"github.com/astaxie/beego"
)

// 基础的控制器，其他控制器在这里继承
type BaseController struct {
	beego.Controller
}

//检测用户是否登录
func (this *BaseController) Prepare() {
	if !IsLogin(this.Ctx) {
		this.Redirect("/manager/login", 302)
		return
	}
}

//检测是否是ajax请求
func (this *BaseController) IsAjax() {
	if !this.Ctx.Input.IsAjax() {
		this.Abort("404")
	}
}

/* End of file  : base.go */
/* Location     : ./controllers/manager/base.go */
