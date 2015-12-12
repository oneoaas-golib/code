package manager

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

// 文章结构体
type ArticleController struct {
	BaseController
}

// 显示文章首页
func (this *ArticleController) Get() {
	//获取类型
	atype := this.Ctx.Input.Param(":type")
	var state []int
	if "trash" != atype {
		state = []int{1}
	} else {
		state = []int{0}
	}

	// 处理分页
	pageSize, err := beego.AppConfig.Int("pagesize")
	if err != nil {
		pageSize = 10
	}
	count, err := models.GetArticleCount(state)
	if err != nil {
		beego.Error(err)
	}
	paginator := pagination.NewPaginator(this.Ctx.Request, pageSize, count)
	this.Data["paginator"] = paginator

	// 查询数据库
	this.Data["Articles"], err = models.GetArticles(paginator.Offset(), pageSize, state)
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsTrash"] = atype == "trash"
	this.Data["Title"] = "管理后台 - 所有文章"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_index_heade.html"
	return
}

// 创建文章
func (this *ArticleController) Create() {
	//显示创建的页面
	if this.Ctx.Input.Method() == "GET" {
		var err error
		this.Data["Title"] = "管理后台 - 创建文章"
		this.Data["Categories"], err = models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/article_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/article_create_heade.html"
		return
	}
	// 处理 post 请求
	title := this.GetString("title")
	category := this.GetString("category")
	content := this.GetString("content")
	if title == "" || category == "" || content == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	err := models.AddArticle(title, category, content)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章添加成功！"}
	}
	this.ServeJson()
	return
}

// 删除文章
func (this *ArticleController) Delete() {
	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}
	err = models.DelArticle(id)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章删除成功！"}
	}
	this.ServeJson()
	return
}

// 修改文章
func (this *ArticleController) Edit() {
	// 显示修改文章的页面
	if this.Ctx.Input.Method() == "GET" {
		id := this.Ctx.Input.Param(":id")
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		this.Data["Categories"], err = models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Article"], err = models.GetArticle(intid)
		if err != nil {
			this.Abort("404")
		}
		this.Data["Title"] = "管理后台 - 修改文章"
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/article_edit.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/article_edit_heade.html"
		return
	}
	// 处理修改文章请求
	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}
	title := this.GetString("title")
	category := this.GetString("category")
	content := this.GetString("content")
	if title == "" || category == "" || content == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	err = models.EditArticle(id, title, category, content)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章修改完成！"}
	}
	this.ServeJson()
	return
}

// 查看一篇文章
func (this *ArticleController) View() {
	id := this.Ctx.Input.Param(":id")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Article"], err = models.GetArticle(intid)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Title"] = "管理后台 - 预览文章"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_view.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_view_heade.html"
	return
}

// 移动到回收站
func (this *ArticleController) Trash() {
	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		return
	}
	//判断操作
	trash := this.GetString("trash")
	if trash == "true" {
		err = models.Trash(id, true)
	} else {
		err = models.Trash(id, false)
	}
	//判断操作结果
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "文章移动到了回收站！"}
	}
	this.ServeJson()
	return
}

/* End of file : article.go */
/* Location : ./controllers/manager/article.go */
