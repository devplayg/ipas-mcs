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

func (c *IpaslogController) LangPrepare() {
	c.addToFrontLang("ipas.start,shock,speeding,proximity")
}

func (c *IpaslogController) Get() {
	filter := c.getFilter()

	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		filter := c.getFilter()
		logs, total, err := models.GetIpaslog(filter)
		c.serveResultJson(logs, total, err, filter.FastPaging)
	} else { // Ajax 외 요청이면 HTML 리턴
		c.Data["filter"] = filter
		c.setTpl("ipaslog.tpl")
	}
}

func (c *IpaslogController) Post() {
	c.Get()
}

func (c *IpaslogController) getFilter() *objs.IpasFilter {

	// 요청값 분류
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
	if filter.StartDate == "" || filter.EndDate == "" {
		t := time.Now()
		filter.StartDate = t.AddDate(0, 0, -7).Format("2006-01-02") + " 00:00"
		filter.EndDate = t.Format("2006-01-02") + " 23:59"
	}

	// 페이징 처리
	if filter.Sort == "" {
		filter.Sort = "date"
	}
	if filter.Order == "" {
		filter.Order = "desc"
	}
	if filter.Limit < 1 {
		filter.Limit = 15
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}


