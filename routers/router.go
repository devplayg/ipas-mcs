package routers

import (
	"github.com/devplayg/ipasm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
