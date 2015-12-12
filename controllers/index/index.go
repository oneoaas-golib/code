package index

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

//博客前台的页面
type IndexController struct {
	beego.Controller
}

//博客首页
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

//查看一篇文章
func (this *IndexController) Posts() {
	id := this.Ctx.Input.Param(":id")
	intid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Article"], err = models.GetArticle(intid)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		this.Abort("404")
	}
	this.Layout = "index/layout.html"
	this.TplNames = "index/view.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "index/view_heade.html"
	return
}

/* End of file 	: index.go */
/* Location 	: ./controllers/index/index.go */
