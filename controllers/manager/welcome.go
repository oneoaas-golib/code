package manager

type WelcomeController struct {
	BaseController
}

func (this *WelcomeController) Get() {
	this.Data["Title"] = "梵响 - 后台首页"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/index.html"
}
