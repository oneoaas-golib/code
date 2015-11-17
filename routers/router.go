package routers

import (
	"code/controllers/manager"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &manager.WelcomeController{})
}
