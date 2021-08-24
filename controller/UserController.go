package controller

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"golearning/models"
)

type UserController struct {
	web.Controller
}

func (this *UserController) Prepare() {
	fmt.Printf("user controller init...")
	this.Data["json"] = map[string]interface{}{"error":"没有权限", "errno":401}
	this.ServeJSON()
	this.StopRun()
}


func (this *UserController) Get() {
	user := models.User{1, "leo", "leozeng"}
	// 处理逻辑
	fmt.Printf("user controller get...")
	this.Data["json"] = &user
	this.ServeJSON()
}

func (this *UserController) Post() {
	fmt.Printf("user controller post...")
	// 处理逻辑
	//this.ServeJSON()
}

func (this *UserController) login()  {
	fmt.Printf("user controller login...")

}