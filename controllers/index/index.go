package index

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	//先处理分页
	pageSize, err := beego.AppConfig.Int("pagesize")
	if err != nil {
		beego.Error(err)
	}
	count, err := models.GetArticleCount([]int{1})
	if err != nil {
		beego.Error(err)
	}
	paginator := pagination.NewPaginator(this.Ctx.Request, pageSize, count)
	this.Data["paginator"] = paginator

	this.Data["Articles"], err = models.GetArticles(paginator.Offset(), pageSize, []int{1})
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Layout = "index/layout.html"
	this.TplNames = "index/index.html"
	return
}

func (this *IndexController) Posts() {
	id := this.Ctx.Input.Param("id")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Article"], err = models.GetArticle(intid)
	if err != nil {
		this.Abort("404")
	}
	this.Layout = "index/layout.html"
	this.TplNames = "index/index.html"
	return
}
