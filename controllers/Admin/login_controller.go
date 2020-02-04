package Admin

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type LoginController struct {
	Ctx iris.Context
}

func (a *LoginController) GetLogin() mvc.View  {
	return mvc.View {
		Layout:"",
		Name:"admin/login.html",
	}
}
