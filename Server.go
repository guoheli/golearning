package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("/static")
	s.Domain("localhost").BindHandler("/", Hello)

	//s.Group("/", func(group *ghttp.RouterGroup) {
		//group.PUT("/save", student.Save)
		//group.GET("/select", student.Select)
		//group.POST("/update", student.Update)
		//group.DELETE("/delete", student.Delete)
	//})
	s.SetPort(8999, 9000)
	s.Run()
}


func Hello(r *ghttp.Request) {
	r.Response.Write("哈喽世界！")
}