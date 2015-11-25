package manager

import (
	"code/models"
	"github.com/astaxie/beego"
)

//文章结构体
type ArticleController struct {
	BaseController
}

//显示文章首页
func (this *ArticleController) Get() {
	var err error
	this.Data["Articles"], err = models.GetArticles("0", 10)
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_index_heade.html"
}

//创建文章
func (this *ArticleController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		var err error
		this.Data["Categories"], err = models.GetCategories("1", 10)
		if err != nil {
			beego.Error(err)
		}
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/article_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/article_create_heade.html"
		return
	}

	title := this.GetString("title")
	category := this.GetString("category")
	content := this.GetString("content")
	state := this.GetString("state")
	if title == "" || category == "" || content == "" || state == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	err := models.AddArticle(title, category, content, state)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章添加成功！"}
	}
	this.ServeJson()
	return
}

func (this *ArticleController) Delete() {
	if !this.Ctx.Input.IsAjax() {
		this.Ctx.WriteString("请求错误！")
		return
	}

	id := this.GetString("id")
	err := models.DelArticle(id)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章删除成功！"}
	}
	this.ServeJson()
	return
}

func (this *ArticleController) Edit() {
	if this.Ctx.Input.Method() == "GET" {
		id := this.Ctx.Input.Param(":id")
		var err error
		this.Data["Categories"], err = models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Article"], err = models.GetArticle(id)
		if err != nil {
			beego.Error(err)
		}
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/article_edit.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/article_edit_heade.html"
		return
	}

	id := this.GetString("id")
	title := this.GetString("title")
	category := this.GetString("category")
	content := this.GetString("content")
	state := this.GetString("state")
	if id == "" || title == "" || category == "" || content == "" || state == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	err := models.EditArticle(id, title, category, content, state)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章修改完成！"}
	}
	this.ServeJson()
	return
}

//查看一篇文章
func (this *ArticleController) View() {
	id := this.Ctx.Input.Param(":id")
	var err error
	this.Data["Article"], err = models.GetArticle(id)
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_view.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_view_heade.html"
	return
}

/* End of file : article.go */
/* Location : ./controllers/manager/article.go */
