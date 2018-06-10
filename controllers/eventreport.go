package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type EventReportController struct {
	baseController
}

func (c *EventReportController) CtrlPrepare() {
	//추가 언어 키워드
	//c.addToFrontLang("ipas.start,shock,speeding,proximity")

	// 권한 부여
	c.grant(objs.User)
}

func (c *EventReportController) getFilter() *objs.ReportFilter {
	filter := objs.ReportFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}
	filter.OrgId, _ = strconv.Atoi(c.Ctx.Input.Param(":orgId"))
	filter.EquipId = c.Ctx.Input.Param(":equipId")
	filter.SinceDays, _ = strconv.Atoi(c.Ctx.Input.Param(":sinceDays"))

	t := time.Now()
	if filter.SinceDays == 0 {
		filter.StartDate = t.Format(objs.DateOnlyFormat) + " 00:00"
		filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
	} else {
		//t.Add(filter.SinceDays * time.Hour * 24)
		//dur := time.Duration(24 * 24 * 60 * filter.SinceDays * time.Second * -1)
		//dur := time.Duration(filter.SinceDays)*time.Second*24 * 24 * 60*-1
		dur := time.Hour * time.Duration(24*filter.SinceDays*-1)
		filter.StartDate = t.Add(dur).Format(objs.DateOnlyFormat) + " 00:00"
		filter.EndDate = t.Format(objs.DateOnlyFormat) + " 23:59"
	}

	return &filter
}
func (c *EventReportController) GetReportData() {
	filter := c.getFilter()
	c.Data["json"] = c.getReport(filter)
	c.ServeJSON()
}

func (c * EventReportController) getTracksAndEvents(filter *objs.ReportFilter) ([]objs.LocTrack, []objs.IpasLog) {
	logFilter := objs.IpasFilter{}
	logFilter.StartDate = filter.StartDate
	logFilter.EndDate = filter.EndDate
	logFilter.OrgId = []int{filter.OrgId}
	logFilter.EquipId = filter.EquipId
	logFilter.FastPaging = "on"
	logFilter.Sort = "date"
	logFilter.Order = "asc"
	logFilter.Limit = 99999

	rows, _, err := models.GetIpaslog(&logFilter, c.member)
	if err != nil {
		log.Error(err)
	}

	tracks := make([]objs.LocTrack, len(rows))
	size := 10 // 최근 10개 이벤트 조회
	events := make([]objs.IpasLog, size)
	for idx, r := range rows {
		if idx <size {
			events[idx] = rows[len(rows)-1-idx]
		}
		tracks[idx] = objs.LocTrack{r.Date, r.Latitude, r.Longitude}
	}
	return tracks, events
}

func (c *EventReportController) getReport(filter *objs.ReportFilter) interface{} {
	m := make(map[string]interface{})

	// IPAS 정보
	ipas, err := models.GetIpas(filter.OrgId, filter.EquipId)
	if err != nil {
		log.Error(err)
	}
	
	// IPAS 정보
	ipas.OrgName, _ = GetOrgGroupName(filter.OrgId, filter.GroupId)
	m["ipas"] = ipas

	// 건수 정보
	m["counts"] = c.getCounts(filter)

	// 트랙 및 이벤트 정보
	m["tracks"], m["events"] = c.getTracksAndEvents(filter)

	return m
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
