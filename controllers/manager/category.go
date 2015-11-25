package manager

import (
	"code/models"
	"github.com/astaxie/beego"
)

//分类的结构体
type CategoryController struct {
	BaseController
}

//显示首页
func (this *CategoryController) Get() {
	var err error
	this.Data["Categories"], err = models.GetCategories("0", "10")
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/category_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/category_index_heade.html"
}

//创建分类
func (this *CategoryController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/category_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/category_create_heade.html"
		return
	}

	name := this.GetString("name")
	desc := this.GetString("description")
	if name == "" || desc == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "分类标题或描述不能为空！"}
		this.ServeJson()
		return
	}

	err := models.AddCategory(name, desc)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "分类添加成功！"}
	}
	this.ServeJson()
	return
}

//删除分类
func (this *CategoryController) Delete() {
	if !this.Ctx.Input.IsAjax() {
		this.Ctx.WriteString("请求错误")
		return
	}
	id := this.GetString("id")
	if id == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "请求错误！"}
		this.ServeJson()
		return
	}
	err := models.DelCategory(id)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "删除成功！"}
	}
	this.ServeJson()
	return
}

//修改分类
func (this *CategoryController) Edit() {
	if this.Ctx.Input.Method() == "GET" {
		id := this.Ctx.Input.Param(":id")
		var err error
		this.Data["Category"], err = models.GetCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/category_edit.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/category_edit_heade.html"
		return
	}

	id := this.GetString("id")
	name := this.GetString("name")
	desc := this.GetString("description")
	if name == "" || desc == "" || id == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "分类名称或描述不能为空！"}
		this.ServeJson()
		return
	}
	err := models.EditCategory(id, name, desc)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "分类修改成功！"}
	}
	this.ServeJson()
	return
}

/* End of file : category.go */
/* Location : ./controllers/manager/category.go */
