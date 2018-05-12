package routers

import (
	"github.com/devplayg/ipas-mcs/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 대시보드
	beego.Router(`/dashboard`, &controllers.DashboardController{}, "get:Display")

	// IPAS 로그
	beego.Router(`/ipaslogs`, &controllers.IpaslogController{})
	beego.Router(`/getIpasLogs`, &controllers.IpaslogController{}, "get:GetLogs")
	beego.Router(`/realtimelogs`, &controllers.IpaslogController{}, "get,post:DisplayRealTimeLogs")
	beego.Router(`/getRealTimeLogs`, &controllers.IpaslogController{}, "get:GetRealTimeLogs")

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
	beego.Router("/assetclass/:class:int/root/:assetId:int", &controllers.AssetController{}, "get:GetChildrenWithRoot")
	//beego.Router("/assets/:assetId:int/children", &controllers.AssetController{}, "Get:GetDescendants")
	beego.Router("/assets/:assetId:int", &controllers.AssetController{}, "Get:GetAsset")
	beego.Router("/assets/:assetId:int", &controllers.AssetController{}, "Patch:UpdateAsset")
	beego.Router("/assets/delete", &controllers.AssetController{}, "Post:RemoveAsset")

	// 사용자 자산
	beego.Router("/userassetclass/:class/children", &controllers.UserassetController{}, "Get:GetChildren")
	beego.Router("/ipasorg/:orgId:int", &controllers.IpaslistController{}, "Get:GetIpasInOrg")
	beego.Router("/ipasgroup/:groupId:int", &controllers.IpaslistController{}, "Get:GetIpasInGroup")
	beego.Router("/ipasgroup/:groupId:int", &controllers.IpaslistController{}, "Patch:UpdateIpasGroup")

	// 통계
	beego.Router(`/stats/:statsType/org/:orgId(-?[\d]+)/group/:groupId(-?[\d]+)`, &controllers.StatsController{}, "Get:GetStats")
	beego.Router(`/stats/:statsType/by/:assetType/org/:orgId(-?[\d]+)/group/:groupId(-?[\d]+)`, &controllers.StatsController{}, "Get:GetStatsBy")
	beego.Router(`/stats/summary/org/:orgId(-?[\d]+)/group/:groupId(-?[\d]+)`, &controllers.StatsController{}, "Get:GetSummary")

	// 환경설정
	beego.Router(`/config`, &controllers.ConfigController{})

	// 로그인
	beego.Router("/", &controllers.LoginController{}, "get:Get")
	beego.Router(`/signin`, &controllers.LoginController{})
	beego.Router(`/signout`, &controllers.LoginController{}, "*:Logout")
	beego.Router(`/signin/:username([\w]+)/salt`, &controllers.LoginController{}, "Get:GetPasswordSalt")
}
