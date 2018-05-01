package controllers

import "github.com/devplayg/ipas-mcs/objs"

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
	// 권한 부여
	c.setTpl("dashboard.tpl")
}