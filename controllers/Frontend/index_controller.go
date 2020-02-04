package Frontend

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	ctx iris.Context
}

func (i IndexController ) Get() mvc.View  {
	return mvc.View{
		Name:"frontend/index.html",
	}
}

func (i *IndexController) GetShop() mvc.View  {
	return mvc.View{
		Name:"frontend/product-details.html",
	}
}

func (i *IndexController) GetList() mvc.View  {
	return mvc.View{
		Name:"frontend/shop.html",
	}
}
