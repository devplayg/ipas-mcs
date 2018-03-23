package routers

import (
	"github.com/devplayg/ipas-mcs/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 대시보드

	// IPAS 로그
	beego.Router(`/ipaslogs`, &controllers.IpaslogController{})

	// Sample
	beego.Router(`/samplelogs`, &controllers.SamplelogController{})

	// 사용자
	beego.Router(`/members`, &controllers.MemberController{})
	beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{})

	// 자산
	beego.Router(`/assets`, &controllers.AssetController{})
	beego.Router("/assetclass/:class:int/descendants/:assetId:int", &controllers.AssetController{}, "get:GetDescendants")
	beego.Router("/assets/:assetId:int/children", &controllers.AssetController{}, "get:GetChildren")
	//beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{}, "Get:GetMemberById")
	//beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{}, "Post:UpdateMember")

	// Post, Delete, Patch, Get(html, json)
	/*
		post /members
		delete /members/1
		patch / members/1
		delete /members/1
	 */

	// 제어

	// 환경설정
	beego.Router(`/config`, &controllers.ConfigController{})

	// 로그인
	beego.Router("/", &controllers.LoginController{}, "get:Get")
	beego.Router(`/signin`, &controllers.LoginController{})
	beego.Router(`/signout`, &controllers.LoginController{}, "*:Logout")
	beego.Router(`/signin/:username([\w]+)/salt`, &controllers.LoginController{}, "Get:GetPasswordSalt")
}
