package main

import (
	"SataCerPlatform/db_mysql"
	_ "SataCerPlatform/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func main() {
  db_mysql.Connect()

  beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

