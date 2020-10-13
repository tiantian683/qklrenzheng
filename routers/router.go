package routers

import (
	"SataCerPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/reg",&controllers.Regcontrollers{})
}
