package manager

type WelcomeController struct {
	BaseController
}

//显示后台首页
func (this *WelcomeController) Get() {
	this.Data["Title"] = "梵响 - 后台首页"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/index.html"
}

/* End of file  : welcome.go */
/* Location     : ./controllers/manager/welcome.go */
