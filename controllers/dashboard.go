package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/astaxie/beego"
)

type DashboardController struct {
	baseController
}

func (c *DashboardController) CtrlPrepare() {
	// 추가 언어 키워드
	c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *DashboardController) Display() {
	c.Data["daumMapKey"] = beego.AppConfig.DefaultString("daummapkey", "IPAS-MCS")

	// 권한 부여
	c.setTpl("dashboard.tpl")
}