package routers

import (
	"github.com/devplayg/ipas-mcs/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 대시보드

	// IPAS 로그
	beego.Router(`/ipaslogs`, &controllers.IpaslogController{})

	// 사용자
	beego.Router(`/members/`, &controllers.MemberController{})

	// 제어

	// 환경설정

	// 로그인
	beego.Router("/", &controllers.LoginController{}, "get:Get")
	beego.Router(`/signin`, &controllers.LoginController{})
	beego.Router(`/signout`, &controllers.LoginController{}, "*:Logout")
	beego.Router(`/signin/:username([\w]+)/salt`, &controllers.LoginController{}, "Get:GetPasswordSalt")
}
