package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
	"sort"
	"time"
	"encoding/json"
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
	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		c.GetLogs()
	} else { // Ajax 외 요청이면 HTML 리턴
		c.Display()
	}
}

func (c *IpaslogController) Post() {
	c.Get()
}

func (c *IpaslogController) Display() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("ipas_log.tpl")
}

func (c *IpaslogController) GetLogs() {
	filter := c.getFilter()
	logs, total, err := models.GetIpaslog(filter, c.member)
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

//func (c *IpaslogController) getLogs() ([]objs.IpasLog, int64){
//	filter := c.getFilter()
//	logs, total, err := models.GetIpaslog(filter, c.member)
//	if err != nil {
//		log.Error(err)
//	}
//
//	// 기관/그룹코드를 이름과 맵핑
//	for idx, a := range logs {
//		logs[idx].No = filter.PagingFilter.Offset + int64(idx) + 1
//		logs[idx].OrgName, logs[idx].GroupName = GetOrgGroupName(a.OrgId, a.GroupId)
//	}
//
//	return logs, total
//}

func (c *IpaslogController) GetMapLogs() {
	filter := c.getFilter()
	filter.Limit = 70
	logs, _, err := models.GetIpaslog(filter, c.member)
	if err != nil {
		log.Error(err)
	}

	// 기관/그룹코드를 이름과 맵핑
	for idx, a := range logs {
		logs[idx].No = filter.PagingFilter.Offset + int64(idx) + 1
		logs[idx].OrgName, logs[idx].GroupName = GetOrgGroupName(a.OrgId, a.GroupId)
	}

	mapLogs := make([]objs.IpasMapLog, len(logs))
	for idx, r := range logs {
		mapLogs[idx].OrgId       = r.OrgId
		mapLogs[idx].EquipId     = r.EquipId
		mapLogs[idx].GroupId     = r.GroupId
		mapLogs[idx].EquipType   = r.EquipType
		mapLogs[idx].Speed       = r.Speed
		mapLogs[idx].Latitude    = FloatToString(r.Latitude)
		mapLogs[idx].Longitude   = FloatToString(r.Longitude)
		mapLogs[idx].OrgName     = r.OrgName
		mapLogs[idx].GroupName   = r.GroupName
		mapLogs[idx].Date        = r.Date
		mapLogs[idx].EventType   = r.EventType
		mapLogs[idx].Targets     = r.Targets
		mapLogs[idx].Distance    = r.Distance
		mapLogs[idx].Label    = "dsfasdf"
	}
	result, err := json.Marshal(mapLogs)
	if err != nil {
		log.Error(err)
	}
	c.Ctx.ResponseWriter.Write([]byte("mapfeed_callback("+ string(result)+")"))
}

func (c *IpaslogController) getFilter() *objs.IpasFilter {

	// 요청값 분류
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 날짜 설정
	if !filter.StatsMode { // 일반적으로 로그를 조회하는 경우
		if filter.StartDate == "" || filter.EndDate == "" {
			t := time.Now().In(c.member.Location)
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

	// 이벤트 맵
	if filter.EventMap == "" {
		filter.EventMap = "on"
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

func (c *IpaslogController) DisplayRealTimeLogs() {
	c.setTpl("realtime_log.tpl")
}

func (c *IpaslogController) GetRealTimeLogs() {
	filter := c.getFilter()

	// 필터 제거
	filter.FastPaging = "on"
	logs, total, err := models.GetIpaslog(filter, c.member)
	if err != nil {
		log.Error(err)
	}

	// 기관/그룹코드를 이름과 맵핑
	for idx, a := range logs {
		logs[idx].OrgName, logs[idx].GroupName = GetOrgGroupName(a.OrgId, a.GroupId)
		logs[idx].DateAgo = humanize.Time(logs[idx].Date)
	}

	c.serveResultJson(logs, total, err, filter.FastPaging)
}

// Echarts
func (c *IpaslogController) GetLogForCharting() { // ipaslog
	filter := c.getFilter()

	// 차트 데이터 초기화
	from, _ := time.ParseInLocation(objs.DefaultDateFormat, filter.StartDate+":00", time.Local)
	to, _ := time.ParseInLocation(objs.DefaultDateFormat, filter.EndDate+":59", time.Local)

	startupChartData := objs.NewTimeLineData(from, to, 3600)
	shockChartData := objs.NewTimeLineData(from, to, 3600)
	speedingChartData := objs.NewTimeLineData(from, to, 3600)
	proximityChartData := objs.NewTimeLineData(from, to, 3600)

	// 통계정보 입력
	rows, _, err := models.GetLogForCharting(filter, c.member)
	if err != nil {
		log.Error(err)
	}
	for _, r := range rows {
		t := r.Date.Unix()

		if r.EventType == objs.StartupEvent {
			if _, ok := startupChartData[r.Date.Unix()]; ok {
				startupChartData[t] += r.Count
			}
		} else if r.EventType == objs.ShockEvent {
			if _, ok := shockChartData[r.Date.Unix()]; ok {
				shockChartData[t] += r.Count
			}
		} else if r.EventType == objs.SpeedingEvent {
			if _, ok := speedingChartData[r.Date.Unix()]; ok {
				speedingChartData[t] += r.Count
			}
		} else if r.EventType == objs.ProximityEvent {
			if _, ok := proximityChartData[r.Date.Unix()]; ok {
				proximityChartData[t] += r.Count
			}
		}
	}

	// 시간 정렬
	timeLines := make(objs.Int64Slice, 0)
	for t := range shockChartData {
		timeLines = append(timeLines, t)
	}
	sort.Sort(timeLines)

	// 차트 데이터 생성
	type timeData map[string][2]int64
	result := map[string][]timeData{
		"startup":   make([]timeData, 0),
		"shock":     make([]timeData, 0),
		"speeding":  make([]timeData, 0),
		"proximity": make([]timeData, 0),
	}
	for _, t := range timeLines {
		result["startup"] = append(result["startup"], timeData{"value": [2]int64{t * 1000, int64(startupChartData[t])}})
		result["shock"] = append(result["shock"], timeData{"value": [2]int64{t * 1000, int64(shockChartData[t])}})
		result["speeding"] = append(result["speeding"], timeData{"value": [2]int64{t * 1000, int64(speedingChartData[t])}})
		result["proximity"] = append(result["proximity"], timeData{"value": [2]int64{t * 1000, int64(proximityChartData[t])}})
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *IpaslogController) DisplayTrend() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("trend.tpl")
}


func (c *IpaslogController) DisplayMap() {
	filter := c.getFilter()
	c.Data["filter"] = filter
	c.setTpl("map.tpl")
}


// For highcharts
//func (c *IpaslogController) GetLogForCharting() { // ipaslog/
//
//	filter := c.getFilter()
//	spew.Dump(filter)
//
//	// 차트 초기화
//	from, _ := time.ParseInLocation(objs.SearchTimeFormat, filter.StartDate, c.member.Location)
//	to, _ := time.ParseInLocation(objs.SearchTimeFormat, filter.EndDate, c.member.Location)
//	shockChartData := objs.NewHighchartsData(from, to, 3600, c.member.Location)
//	speedingChartData := objs.NewHighchartsData(from, to, 3600, c.member.Location)
//	proximityChartData := objs.NewHighchartsData(from, to, 3600, c.member.Location)
//
//	// 통계정보 입력
//	rows, _, err := models.GetLogForCharting(filter, c.member)
//	if err != nil {
//		log.Error(err)
//	}
//	//loc, _ := time.LoadLocation(c.member.Location)
//	for _, r := range rows {
//		t := r.Date.In(c.member.Location).Unix()
//
//		if r.EventType == objs.ShockEvent {
//			if _, ok := shockChartData.TimeMap[r.Date.Unix()]; ok {
//				shockChartData.TimeMap[t] += r.Count
//			}
//		} else if r.EventType == objs.SpeedingEvent {
//			if _, ok := speedingChartData.TimeMap[r.Date.Unix()]; ok {
//				speedingChartData.TimeMap[t] += r.Count
//			}
//		} else if r.EventType == objs.ProximityEvent {
//			if _, ok := proximityChartData.TimeMap[r.Date.Unix()]; ok {
//				proximityChartData.TimeMap[t] += r.Count
//			}
//		}
//	}
//
//	shockChartData.Sort()
//	speedingChartData.Sort()
//	proximityChartData.Sort()
//
//	c.Data["json"] = map[string]*objs.ChartData{
//		"shock":     shockChartData,
//		"speeding":  speedingChartData,
//		"proximity": proximityChartData,
//	}
//	c.ServeJSON()
//}



func FloatToString(f float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(f), 'f', 6, 64)
}