package routers

import (
	"github.com/devplayg/ipas-mcs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 로그인
    beego.Router("/", &controllers.LoginController{}, "get:Get")
	beego.Router(`/signin`, &controllers.LoginController{})
	beego.Router(`/signout`, &controllers.LoginController{}, "*:Logout")
	beego.Router(`/signin/:username([\w]+)/salt`, &controllers.LoginController{}, "Get:GetPasswordSalt")
    ///login/member/salt

    // 대시보드

    // IPAS 로그
	beego.Router(`/ipaslogs`, &controllers.IpaslogController{})

	// 사용자
	beego.Router(`/members/`, &controllers.MemberController{})
	//beego.Router(`/members/, &controllers.IpaslogController{}, "get:GetLogs")

    // 제어어

    // 환경설정
}
