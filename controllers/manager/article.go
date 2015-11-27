package manager

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"os"
	"strconv"
	"time"
)

// 文章结构体
type ArticleController struct {
	BaseController
}

// 显示文章首页
func (this *ArticleController) Get() {
	// 处理分页
	pageSize := 20
	count, err := models.GetArticleCount()
	if err != nil {
		beego.Error(err)
	}
	paginator := pagination.NewPaginator(this.Ctx.Request, pageSize, count)
	this.Data["paginator"] = paginator
	// 查询数据库
	this.Data["Articles"], err = models.GetArticles(paginator.Offset(), pageSize)
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_index_heade.html"
}

// 创建文章
func (this *ArticleController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		categories, err := models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Categories"] = categories
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
	if !this.Ctx.Input.IsAjax() {
		this.Ctx.WriteString("请求错误！")
		return
	}

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
			beego.Error(err)
		}
		this.Data["Categories"], err = models.GetAllCategories()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Article"], err = models.GetArticle(intid)
		if err != nil {
			beego.Error(err)
		}
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
		beego.Error(err)
	}
	this.Data["Article"], err = models.GetArticle(intid)
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/article_view.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/article_view_heade.html"
	return
}

// 移动上传的文件
func (this *ArticleController) MoveUploadFile() {
	if !this.Ctx.Input.IsAjax() {
		this.Ctx.WriteString("请求错误！")
		return
	}
	filename := this.GetString("filename")
	datePath := time.Now().Format("2006/01")
	dirPath := "./upload/" + datePath
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		beego.Error(err)
	}
	err = os.Rename("./tmp/"+filename, dirPath+"/"+filename)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": "success",
			"data": map[string]string{
				"filepath": "/upload/" + datePath + "/" + filename,
			},
		}
	}
	this.ServeJson()
	return
}

func (this *ArticleController) RemoveToTrash() {}

func (this *ArticleController) ReturnFromTrash() {}

/* End of file : article.go */
/* Location : ./controllers/manager/article.go */
