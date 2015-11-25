package manager

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Post() {
	file, filehead, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
}
