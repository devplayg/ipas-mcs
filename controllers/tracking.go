package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"time"
)

type TrackingController struct {
	baseController
}

func (c *TrackingController) CtrlPrepare() {
	// 추가 언어 키워드
	//c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *TrackingController) Display() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("tracking.tpl")
}

func (c *TrackingController) getFilter() *objs.IpasTrackingFilter {

	// 요청값 분류
	filter := objs.IpasTrackingFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
	if filter.StartDate == "" || filter.EndDate == "" {
		t := time.Now()
		//filter.StartDate = t.AddDate(0, 0, -7).Format(objs.DateOnlyFormat) + " 00:00"
		filter.StartDate = t.Add(-86400*7*time.Second).Format(objs.DateOnlyFormat) + " 00:00"
		filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"

	} else {
		filter.StartDate = filter.StartDate + " 00:00"
		filter.EndDate = filter.StartDate + " 23:59"
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
