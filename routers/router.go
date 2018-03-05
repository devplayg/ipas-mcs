package routers

import (
	"github.com/devplayg/ipas-mcs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 로그인
    beego.Router("/", &controllers.LoginController{})

    // 대시보드

    // 로그

    // 제어어

    // 환경설정
}
