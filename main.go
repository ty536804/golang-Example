package main

import (
	"context"
	"express/controllers/Admin"
	"express/controllers/Frontend"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main()  {
	app :=iris.New()
	app.Logger().SetLevel("debug")

	template := iris.HTML("./resources/views", ".html").Reload(true)
		//Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)

	app.StaticWeb("/assets", "./resources/assets")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().
			GetStringDefault("message","访问的页面出错了~！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	//后端
	adminParty := app.Party("/admin")
	admin := mvc.New(adminParty)
	admin.Register(ctx)
	admin.Handle(new(Admin.LoginController))

	//前端
	indexParty := app.Party("/")
	index := mvc.New(indexParty)
	index.Register(ctx)
	index.Handle(new(Frontend.IndexController))

	app.Run(iris.Addr("0.0.0.0:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
