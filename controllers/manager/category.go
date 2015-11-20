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
	this.TplNames = "manager/category_index.html"
	var err error
	this.Data["Categories"], err = models.GetCategories("0", "10")
	if err != nil {
		beego.Error(err)
	}
}

//创建分类
func (this *CategoryController) Create() {
	if this.Ctx.Input.Method() == "GET" {
		this.TplNames = "manager/category_create.html"
		this.Layout = "manager/layout.html"
		return
	}

	name := this.GetString("name")
	if name == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "分类标题不能为空！"}
		this.ServeJson()
		return
	}

	err := models.AddCategory(name)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "分类添加成功！"}
	}
	this.ServeJson()
	return
}

//删除分类
func (this *CategoryController) Delete() {}

//修改分类
func (this *CategoryController) Edit() {}

/* End of file : category.go */
/* Location : ./controllers/manager/category.go */
