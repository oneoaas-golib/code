package manager

import (
	"code/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
	var err error
	this.Data["Articles"], err = models.GetArticles("0", "10")
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_index_heade.html"
	this.LayoutSections["Scripts"] = ""
}
