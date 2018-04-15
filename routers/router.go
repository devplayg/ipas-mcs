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
	beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{}, "Patch:UpdateMember")
	beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{}, "Delete:DeleteMember")
	beego.Router(`/members/:memberId([\d]+)`, &controllers.MemberController{}, "Get:GetMemberById")
	beego.Router(`/members/:memberId([\d]+)/acl`, &controllers.MemberController{}, "Get:GetMemberAcl")
	beego.Router(`/members/:memberId([\d]+)/acl`, &controllers.MemberController{}, "Patch:UpdateMemberAcl")

	// 자산
	beego.Router("/assets", &controllers.AssetController{})
	beego.Router("/assetclass/:class:int/root/:assetId:int", &controllers.AssetController{}, "get:GetDescendantsWithRoot")
	//beego.Router("/assets/:assetId:int/children", &controllers.AssetController{}, "get:GetChildren")
	beego.Router("/assets/:assetId:int", &controllers.AssetController{}, "Get:GetAsset")
	beego.Router("/assets/:assetId:int", &controllers.AssetController{}, "Patch:UpdateAsset")
	beego.Router("/assets/delete", &controllers.AssetController{}, "Post:RemoveAsset")

	beego.Router("/ipasorg/:orgId:int", &controllers.IpaslistController{}, "Get:GetIpasInOrg")
	beego.Router("/ipasgroup/:groupId:int", &controllers.IpaslistController{}, "Get:GetIpasInGroup")
	beego.Router("/ipasgroup/:groupId:int", &controllers.IpaslistController{}, "Patch:UpdateIpasGroup")

	// users/me
	// /ipaslist, /ipasorg/1/group/1
	// ipasgroup
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
