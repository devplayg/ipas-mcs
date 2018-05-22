package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/devplayg/ipas-mcs/models"
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
	filter := c.getFilter()
	c.Data["filter"] = filter

	c.setTpl("dashboard.tpl")
}

func (c *DashboardController) DisplayDarkboard() {
	filter := c.getFilter()
	c.Data["filter"] = filter

	c.setTpl("darkboard.tpl")
}

func (c *DashboardController) DisplayDetailboard() {
	filter := c.getFilter()
	c.Data["filter"] = filter

	c.setTpl("detailboard.tpl")
}

func (c *DashboardController) getFilter() *objs.IpasFilter {

	// 요청값 분류
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
	if filter.StartDate == "" || filter.EndDate == "" {
		config, err := models.GetSystemConfig("stats", "last_updated")
		if err != nil {
			log.Error(err)
		}
		var t time.Time
		if len(config) == 1 {
			t, _ = time.Parse(objs.DefaultDateFormat, config[0].ValueS)
		} else {
			t = time.Now()
		}
		filter.StartDate = t.Format("2006-01-02")
		filter.EndDate = t.Format("2006-01-02")
	}

	return &filter
}
