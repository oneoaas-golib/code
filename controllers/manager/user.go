package manager

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

//用户控制器
type UserController struct {
	BaseController
}

//显示用户首页
func (this *UserController) Get() {
	pagesize, err := beego.AppConfig.Int("pagesize")
	if err != nil {
		beego.Error(err)
	}
	count, err := models.GetUserCount()
	if err != nil {
		beego.Error(err)
	}
	paginator := pagination.NewPaginator(this.Ctx.Request, pagesize, count)
	this.Data["paginator"] = paginator
	this.Data["Users"], err = models.GetUsers(paginator.Offset(), pagesize)
	if err != nil {
		beego.Error(err)
	}

	this.Layout = "manager/layout.html"
	this.TplNames = "manager/user_index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "manager/user_index_heade.html"
	this.LayoutSections["Pagination"] = "manager/pagination.html"
	if err != nil {
		beego.Error(err)
	}
}

//创建用户
func (this *UserController) Create() {

	if this.Ctx.Input.Method() == "GET" {
		this.Layout = "manager/lauout.html"
		this.TplNames = "manager/user_create.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/user_create_heade.html"
		return
	}

	name := this.GetString("name")
	passone := this.GetString("passone")
	passtwo := this.GetString("passtwo")
	if name == "" || passone == "" || passtwo == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	if passone != passtwo {
		this.Data["json"] = map[string]string{"code": "error", "info": "两次填写的密码不一致！"}
		this.ServeJson()
		return
	}

	err := models.AddUser(name, passone)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "用户添加成功！"}
	}
	this.ServeJson()
	return
}

// 修改用户
func (this *UserController) Edit() {
	if this.Ctx.Input.Method() == "GET" {
		id := this.Ctx.Input.Param(":id")
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["User"], err = models.GetUser(intid)
		if err != nil {
			beego.Error(err)
		}
		this.Layout = "manager/layout.html"
		this.TplNames = "manager/user_edit.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = "manager/user_edit_heade.html"
		return
	}

	id, err := this.GetInt64("id")
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
		this.ServeJson()
		return
	}

	username := this.GetString("username")
	passone := this.GetString("passone")
	passtwo := this.GetString("passtwo")
	if username == "" || passone == "" || passtwo == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "必填选项不能为空！"}
		this.ServeJson()
		return
	}
	if passone != passtwo {
		this.Data["json"] = map[string]string{"code": "error", "info": "两次输入的密码不一致！"}
		this.ServeJson()
		return
	}

	err = models.EditUser(id, username, passone)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "用户修改成功！"}
	}
	this.ServeJson()
	return
}

// 删除用户
func (this *UserController) Delete() {
	this.IsAjax()

	count, err := models.GetUserCount()
	if err != nil {
		beego.Error(err)
	}
	if count == 1 {
		this.Data["json"] = map[string]string{"code": "error", "info": "必须保留一个用户！"}
		this.ServeJson()
		return
	}

	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error(err)
	}
	err = models.DelUser(id)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.Data["json"] = map[string]string{"code": "success", "info": "用户删除成功！"}
	}
	this.ServeJson()
	return
}

/* End of file : user.go */
/* Location : ./controllers/manager/user.go */
