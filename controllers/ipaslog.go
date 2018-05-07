package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-mcs/models"
	"time"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type IpaslogController struct {
	baseController
}

func (c *IpaslogController) CtrlPrepare() {
	// 추가 언어 키워드
	c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *IpaslogController) Get() {
	filter := c.getFilter()

	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		logs, total, err := models.GetIpaslog(filter, c.member)

		// 기관/그룹코드를 이름과 맵핑
		for idx, a := range logs {
			if v, ok:= assetMap.Load(a.OrgId); ok {
				logs[idx].OrgName = v.(objs.Asset).Name
			} else {
				logs[idx].OrgName = strconv.Itoa(a.OrgId)
			}
			if v, ok:= assetMap.Load(a.GroupId); ok {
				logs[idx].GroupName = v.(objs.Asset).Name
			} else {
				logs[idx].GroupName = strconv.Itoa(a.GroupId)
			}
		}

		c.serveResultJson(logs, total, err, filter.FastPaging)
	} else { // Ajax 외 요청이면 HTML 리턴
		c.Data["filter"] = filter
		//c.Data["daumMapKey"] = beego.AppConfig.DefaultString("daummapkey", "IPAS-MCS")
		c.setTpl("ipas_logs.tpl")
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
		filter.Limit = 24
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}


func (c *IpaslogController) DisplayRealTimeLogs() {
	//c.Data["daummap"] = beego.AppConfig.DefaultString("daummap", "IPAS-MCS")
	c.setTpl("realtime_logs.tpl")
}

func (c *IpaslogController) GetRealTimeLogs() {
	filter := c.getFilter()

	// 필터 제거
	t := time.Now()
	filter.StartDate = t.Format("2006-01-02") + " 00:00"
	filter.EndDate = t.Format("2006-01-02") + " 23:59"
	filter.FastPaging = "on"
	filter.Limit = 9

	logs, total, err := models.GetIpaslog(filter, c.member)

	// 기관/그룹코드를 이름과 맵핑
	for idx, a := range logs {
		if v, ok:= assetMap.Load(a.OrgId); ok {
			logs[idx].OrgName = v.(objs.Asset).Name
		} else {
			logs[idx].OrgName = strconv.Itoa(a.OrgId)
		}
		if v, ok:= assetMap.Load(a.GroupId); ok {
			logs[idx].GroupName = v.(objs.Asset).Name
		} else {
			logs[idx].GroupName = strconv.Itoa(a.GroupId)
		}
	}

	c.serveResultJson(logs, total, err, filter.FastPaging)
}