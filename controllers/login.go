package controllers

import "github.com/astaxie/beego"

type Logincontrollers struct {
beego.Controller
}

func (r *Logincontrollers ) Get()  {
	r.TplName = "login.html"
}
