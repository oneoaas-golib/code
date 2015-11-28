package manager

import (
	"code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils/captcha"
)

var capt *captcha.Captcha

type LoginController struct {
	beego.Controller
}

//session存的用户的key
const SESSION_KEY = "session_key"

//初始化的函数，设置验证码
func init() {
	store := cache.NewMemoryCache()
	capt = captcha.NewWithFilter("/captcha/", store)
	capt.StdWidth = 120
	capt.StdHeight = 40
	capt.ChallengeNums = 4
}

//在所有方法前面调用，如果登录既不能显示该页面了
func (this *LoginController) Prepare() {
	if IsLogin(this.Ctx) {
		this.Redirect("/manager", 301)
		return
	}
}

func (this *LoginController) Get() {
	//添加初始用户
	// err := models.AddAdmin()
	// if err != nil {
	// 	beego.Error(err)
	// }
	this.TplNames = "manager/login.html"
}

//处理提交的请求
func (this *LoginController) Post() {
	this.IsAjax()

	//先判断验证码是否正确
	if !capt.VerifyReq(this.Ctx.Request) {
		this.Data["json"] = map[string]string{"code": "error", "info": "验证码错误！"}
		this.ServeJson()
		return
	}
	//判断用户名密码是否为空
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" || password == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "用户名或密码为空！"}
		this.ServeJson()
		return
	}
	//判断用户名或密码是否正确
	err := models.CheckLogin(username, password)
	if err != nil {
		this.Data["json"] = map[string]string{"code": "error", "info": err.Error()}
	} else {
		this.SetSession(SESSION_KEY, username)
		this.Data["json"] = map[string]string{"code": "success", "info": "登录成功！"}
	}
	this.ServeJson()
	return
}

//处理用户退出的操作
func (this *LoginController) Logout() {
	this.DestroySession()
	this.Redirect("/manager/login", 301)
	return
}

//判断用户是否登录
func IsLogin(ctx *context.Context) bool {
	if ctx.Input.Session(SESSION_KEY) == nil {
		return false
	}
	return true
}
