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

func (this *ArticleController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/article_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/article_create_heade.html"
		var err error
		this.Data["Categories"], err = models.GetCategories("0", "10")
		if err != nil {
			beego.Error(err)
		}
		return
	}
}

func (this *ArticleController) Delete() {}

func (this *ArticleController) Edit() {}

func (this *ArticleController) View() {}

/* End of file : article.go */
/* Location : ./controllers/manager/article.go */
