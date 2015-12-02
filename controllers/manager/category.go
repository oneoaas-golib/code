package manager

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

//分类的结构体
type CategoryController struct {
	BaseController
}

//显示首页
func (this *CategoryController) Get() {
	//处理分页
	pageSize, err := beego.AppConfig.Int("pagesize")
	if err != nil {
		pageSize = 10
	}
	count, err := models.GetCategoryCount()
	if err != nil {
		this.Abort("404")
	}
	paginator := pagination.NewPaginator(this.Ctx.Request, pageSize, count)
	this.Data["paginator"] = paginator
	//查询数据
	this.Data["Categories"], err = models.GetCategories(paginator.Offset(), pageSize)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Title"] = "管理后台 - 所有分类"
	this.Layout = "manager/layout.html"
	this.TplNames = "manager/category_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/category_index_heade.html"
}

//创建分类
func (this *CategoryController) Create() {
	//显示创建分类的页面
	if this.Ctx.Input.Method() == "GET" {
		this.Data["Title"] = "管理后台 - 创建分类"
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/category_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/category_create_heade.html"
		return
	}
	//处理创建分类的请求
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
	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}
	err = models.DelCategory(id)
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
	//显示修改的页面
	if this.Ctx.Input.Method() == "GET" {
		id := this.Ctx.Input.Param(":id")
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		this.Data["Category"], err = models.GetCategory(id)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		this.Data["Title"] = "管理后台 - 修改分类"
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/category_edit.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/category_edit_heade.html"
		return
	}
	//处理修改的请求
	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}
	name := this.GetString("name")
	desc := this.GetString("description")
	if name == "" || desc == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "分类名称或描述不能为空！"}
		this.ServeJson()
		return
	}
	err = models.EditCategory(id, name, desc)
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
