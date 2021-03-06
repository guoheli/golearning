package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 注册请求数据结构
type RegisterReq struct {
	Name  string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Pass  string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Pass2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"`
}

// 注册返回数据结构
type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func main0() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/register", func(r *ghttp.Request) {
			var req *RegisterReq
			if err := r.Parse(&req); err != nil {
				r.Response.WriteJsonExit(RegisterRes{
					Code:  1,
					Error: err.Error(),
				})
			}
			// ...
			r.Response.WriteJsonExit(RegisterRes{
				Data: req,
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}