package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-mcs/models"
	"time"
	log "github.com/sirupsen/logrus"
)

type IpaslogController struct {
	baseController
}

func (c *IpaslogController) Get() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("ipaslog.tpl")
}

func (c *IpaslogController) Post() {
	c.Get()
}

func (c *IpaslogController) getFilter() *objs.IpasFilter {

	// 요청값 분류
	log.Debugf("### Limit: %s", c.GetString("limit"))
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}
	log.Debugf("### Parsed limit: %d", filter.Limit)

	// 날짜 설정
	if filter.StartDate == "" || filter.EndDate == "" {
		t := time.Now()
		filter.StartDate = t.AddDate(0, 0, -7).Format("2006-01-02") + " 00:00"
		filter.EndDate = t.Format("2006-01-02") + " 23:59"
	}

	// Paging
	if filter.Sort == "" {
		filter.Sort = "date"
	} else {
		//filter.Sort = libs.CamelToUnderscore(filter.Sort) // Change 'SrcIp' to 'Src_Ip'
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "on"
	}

	if filter.Order == "" {
		filter.Order = "desc"
	}

	if filter.Limit < 1 {
		filter.Limit = 5
	}

	return &filter
}

func (c *IpaslogController) GetLogs() {

	filter := c.getFilter()
	logs, _, err := models.GetIpaslog(filter)
	if err != nil {

	}
	c.Data["json"] = logs
	c.ServeJSON()
}
