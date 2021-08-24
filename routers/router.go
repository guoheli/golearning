package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"golearning/controller"
)

func init() {
	//web.InsertFilter("/", web.BeforeRouter, filterFunc)
	//web.AutoRouter(&controller.UserController{})
	web.Router("/", &controller.UserController{}, "get:get;post:post;get:login")
}

//检查session，此函数由正则匹配的路由，在寻找路由之前执行
func filterFunc(ctx *context.Context) {
	if name := ctx.Input.Session("username"); name == nil {
		ctx.Redirect(302, "/")
		return
	}

}
