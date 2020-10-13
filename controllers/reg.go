package controllers

import (
	"SataCerPlatform/db_mysql"
	"SataCerPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Regcontrollers struct {
	beego.Controller
}

func (r *Regcontrollers) Get() {
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("错误")
		return
	}

	row,err := db_mysql.AddUser(user)
	if err != nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("注册失败，重试")
		return
	}
	fmt.Println(row)

	if row != -1{
		r.TplName="login.html"
		return
	}else {
		fmt.Println("错误，重试")
		return
	}

}