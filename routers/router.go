package routers

import (
	"code/controllers/index"
	"code/controllers/manager"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{})

	beego.Router("/manager", &manager.WelcomeController{})
	beego.Router("/manager/login", &manager.LoginController{})
	beego.Router("/manager/logout", &manager.LoginController{}, "get:Logout")

	beego.Router("/manager/article", &manager.ArticleController{})
}
