package controllers

import (
	"time"
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-mcs/models"
	log "github.com/sirupsen/logrus"
)

type StatuslogController struct {
	baseController
}

func (c *StatuslogController) CtrlPrepare() {
	// 추가 언어 키워드
	//c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *StatuslogController) Get() {
	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		c.GetLogs()
	} else { // Ajax 외 요청이면 HTML 리턴
		c.Display()
	}
}

func (c *StatuslogController) Post() {
	c.Get()
}

func (c *StatuslogController) Display() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("status_log.tpl")
}

func (c *StatuslogController) GetLogs() {
	filter := c.getFilter()
	logs, total, err := models.GetIpasStatusLog(filter, c.member)
	if err != nil {
		log.Error(err)
	}

	// 기관/그룹코드를 이름과 맵핑
	for idx, a := range logs {
		logs[idx].No = filter.PagingFilter.Offset + int64(idx) + 1
		logs[idx].OrgName, logs[idx].GroupName = GetOrgGroupName(a.OrgId, a.GroupId)
	}

	c.serveResultJson(logs, total, err, filter.FastPaging)
}

func (c *StatuslogController) getFilter() *objs.IpasFilter {

	// 요청값 분류
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정

	if !filter.StatsMode { // 일반적으로 로그를 조회하는 경우
		if filter.StartDate == "" || filter.EndDate == "" {
			t := time.Now()
			//filter.StartDate = t.AddDate(0, 0, -7).Format(objs.DateOnlyFormat) + " 00:00"
			//filter.StartDate = t.Add(-86400*7*time.Second).Format(objs.DateOnlyFormat) + " 00:00"
			filter.StartDate = t.Format(objs.DateOnlyFormat) + " 00:00"
			filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
		}

	} else { // 통계 근거로그를 조회하는 경우
		if len(filter.StartDate) > 0 || len(filter.EndDate) > 0 {
			filter.StartDate = filter.StartDate + " 00:00"
			filter.EndDate = filter.EndDate + " 23:59"
		} else if len(filter.StartDate) > 0 { // 특정 지정 날짜에 대한 데이터 조회 시
			filter.StartDate = filter.StartDate + " 00:00"
			filter.EndDate = filter.StartDate + " 23:59"
		} else { // 최종 통계산출 날짜에 대한 데이터 조회 시
			rs, err := models.GetSystemConfig("stats", "last_updated")
			if err != nil {
				log.Error(err)
			}
			if len(rs) == 1 {
				filter.StartDate = rs[0].ValueS[0:10] + " 00:00"
				filter.EndDate = rs[0].ValueS[0:10] + " 23:59"
			} else {
				t := time.Now()
				filter.StartDate = t.Format(objs.DateOnlyFormat) + " 00:00"
				filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
			}
		}
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