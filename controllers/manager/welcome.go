package manager

import (
	"code/models"
	"github.com/astaxie/beego"
)

type WelcomeController struct {
	BaseController
}

//显示后台首页
func (this *WelcomeController) Get() {
	var err error
	this.Data["ArticleCount"], err = models.GetArticleCount([]int{0, 1})
	if err != nil {
		beego.Error(err)
	}
	this.Data["CategoryCount"], err = models.GetCategoryCount()
	if err != nil {
		beego.Error(err)
	}
	this.Data["UserCount"], err = models.GetUserCount()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Title"] = "梵响 - 后台首页"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/index.html"
}

/* End of file  : welcome.go */
/* Location     : ./controllers/manager/welcome.go */
