package main

import "github.com/beego/beego/v2/server/web"

func main()  {
	web.SetStaticPath("/m3u8", "static/m3u8")
	web.Run()
}