package manager

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	// if !IsLogin(this.Ctx) {
	// 	this.Redirect("/manager/login", 301)
	// 	return
	// }
}

/* End of file : base.go */
/* Location : ./controllers/manager/base.go */
