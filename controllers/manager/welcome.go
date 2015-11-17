package manager

type WelcomeController struct {
	BaseController
}

func (this *WelcomeController) Get() {
	this.Ctx.WriteString("hello world")
}
