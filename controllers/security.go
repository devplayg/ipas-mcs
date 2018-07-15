package controllers

import (
	"time"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"github.com/devplayg/ipas-mcs/models"
)

type SecurityController struct {
	baseController
}

func (c *SecurityController) DisplaySecurityLogs() {
	filter := c.getFilter()

	if c.IsAjax() {
		logs, total, err := models.GetSecurityLogs(*filter)
		if err != nil {
			log.Error(err)
		}
		c.serveResultJson(logs, total, err, filter.FastPaging)
	} else {
		c.Data["filter"] = filter
		c.setTpl("security_log.tpl")
	}
}

//func (c *SecurityController) GetSecurityLogs() {
//	filter := c.getFilter()
//	logs, total, err := models.GetSecurityLogs(*filter)
//	c.serveResultJson(logs, total, err, filter.FastPaging)
//}

func (c *SecurityController) getFilter() *objs.SecurityLogFilter {

	// 요청값 분류
	filter := objs.SecurityLogFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
		if filter.StartDate == "" || filter.EndDate == "" {
			t := time.Now()
			filter.StartDate = t.AddDate(0, 0, -7).Format(objs.DateOnlyFormat) + " 00:00"
			filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
		}

	// 페이징 처리
	if filter.Sort == "" {
		filter.Sort = "date"
	}
	if filter.Order == "" {
		filter.Order = "desc"
	}
	if filter.Limit < 1 {
		filter.Limit = 20
	}
	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}
	return &filter
}