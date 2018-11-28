package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"fmt"
)

type UserController struct {
	Ctx iris.Context
}

func (_ *UserController) Get() mvc.View {
	return mvc.View{
		Name:   "user/register.html",
		Layout: "shared/layout.html",
		Data:   iris.Map{"Title": "User Index"},
	}
}
func (my *UserController) PostAdd() string {
	username := my.Ctx.Params().Get("username")
	password := my.Ctx.Params().Get("password")

	fmt.Println(username, password)
}
