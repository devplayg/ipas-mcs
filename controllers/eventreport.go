package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
	"github.com/davecgh/go-spew/spew"
)

type EventReportController struct {
	baseController
}

func (c *EventReportController) CtrlPrepare() {
	// 권한 부여
	c.grant(objs.User)
}

func (c *EventReportController) GetReportData() {
	filter := c.getFilter()
	c.Data["json"] = c.getReport(filter)
	c.ServeJSON()
}

func (c *EventReportController) getReport(filter *objs.ReportFilter) interface{} {
	spew.Dump(filter)
	m := make(map[string]interface{})

	// 기간
	m["date"] = map[string]string{
		"from":  filter.StartDate,
		"to":    filter.EndDate,
		"today": time.Now().Format(time.RFC3339),
	}

	// IPAS 정보
	ipas, err := models.GetIpas(filter.OrgId, filter.EquipId)
	if err != nil {
		log.Error(err)
	}
	ipas.OrgName, _ = GetOrgGroupName(filter.OrgId, filter.GroupId)
	m["ipas"] = ipas

	// 건수 정보
	m["counts"] = c.getCounts(filter)

	// 이벤트 정보
	m["events"] = c.getEvents(filter)

	// 상태 정보
	m["status"] = c.getTracks(filter)

	return m
}

func (c *EventReportController) getFilter() *objs.ReportFilter {
	filter := objs.ReportFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}
	filter.OrgId, _ = strconv.Atoi(c.Ctx.Input.Param(":orgId"))
	filter.EquipId = c.Ctx.Input.Param(":equipId")

	t, err := time.Parse(time.RFC3339, filter.Date)
	if err != nil {
		t = time.Now()
	}

	if filter.PastDays == 0 {
		filter.StartDate = t.Format(objs.DateOnlyFormat) + " 00:00"
		filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
	} else {
		dur := time.Hour * time.Duration(24*filter.PastDays*-1)
		filter.StartDate = t.Add(dur).Format(objs.DateOnlyFormat) + " 00:00"
		filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
	}

	return &filter
}

func (c *EventReportController) getEvents(filter *objs.ReportFilter) []objs.IpasLog {
	logFilter := objs.IpasFilter{}
	logFilter.StartDate = filter.StartDate
	logFilter.EndDate = filter.EndDate
	logFilter.OrgId = []int{filter.OrgId}
	logFilter.EquipId = filter.EquipId
	logFilter.FastPaging = "on"
	logFilter.Sort = "date"
	logFilter.Order = "desc"
	logFilter.Limit = 9999

	rows, _, err := models.GetIpaslog(&logFilter, c.member)
	if err != nil {
		log.Error(err)
	}

	return rows
}

func (c *EventReportController) getTracks(filter *objs.ReportFilter) []objs.IpasLog {
	logFilter := objs.IpasFilter{}
	logFilter.StartDate = filter.StartDate
	logFilter.EndDate = filter.EndDate
	logFilter.OrgId = []int{filter.OrgId}
	logFilter.EquipId = filter.EquipId
	logFilter.FastPaging = "on"
	logFilter.Sort = "date"
	logFilter.Order = "desc"
	logFilter.Limit = 9999

	rows, _, err := models.GetIpasStatusLog(&logFilter, c.member)
	if err != nil {
		log.Error(err)
	}

	return rows
}

func (c *EventReportController) getCounts(filter *objs.ReportFilter) map[string]int {
	// 추이정보
	statsFilter := objs.StatsFilter{
		StartDate: filter.StartDate,
		EndDate:   filter.EndDate,
		OrgId:     filter.OrgId,
		EquipIp:   filter.EquipId,
	}

	statsFilter.StatsType = "equip_trend"
	rows, err := models.GetEquipStats(c.member, statsFilter)
	if err != nil {
		log.Error(err)
	}
	counts := map[int]int{
		objs.StartupEvent:   0,
		objs.ShockEvent:     0,
		objs.SpeedingEvent:  0,
		objs.ProximityEvent: 0,
	}
	var n int
	for _, r := range rows {
		d := strings.SplitN(r.Data, ",", 4)

		n, _ = strconv.Atoi(d[0])
		counts[objs.StartupEvent] += n

		n, _ = strconv.Atoi(d[1])
		counts[objs.ShockEvent] += n

		n, _ = strconv.Atoi(d[2])
		counts[objs.SpeedingEvent] += n

		n, _ = strconv.Atoi(d[3])
		counts[objs.ProximityEvent] += n
	}
	m := map[string]int{
		"startup":   counts[objs.StartupEvent],
		"shock":     counts[objs.ShockEvent],
		"speeding":  counts[objs.SpeedingEvent],
		"proximity": counts[objs.ProximityEvent],
		"activated": 0,
	}

	// 활성화 정보 조회
	statsFilter.StatsType = "activated_equip"
	rows, err = models.GetEquipStats(c.member, statsFilter)
	if err != nil {
		log.Error(err)
	}
	for _, r := range rows {
		m["activated"] += r.Count
	}

	return m
}
