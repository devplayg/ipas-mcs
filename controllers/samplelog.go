package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"time"
)

type SamplelogController struct {
	baseController
}

func (c *SamplelogController) Get() {
	filter := c.getFilter()

	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		filter := c.getFilter()
		logs, total, err := models.GetSamplelog(filter)
		c.serveResultJson(logs, total, err, filter.FastPaging)
	} else { // Ajax 외 요청이면 HTML 리턴
		c.Data["filter"] = filter
		c.setTpl("samplelog.tpl")
	}
}

func (c *SamplelogController) Post() {
	c.Get()
}

func (c *SamplelogController) getFilter() *objs.SampleFilter {

	// 요청값 분류
	filter := objs.SampleFilter{}
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
		filter.Limit = 10
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}
