package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-server"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type node struct {
	Name    string   `json:"name"`
	Size    int      `json:"size"`
	Imports []string `json:"imports"`
}

type StatsController struct {
	baseController
}

func (c *StatsController) CtrlPrepare() {
	// 권한 부여
	c.grant(objs.User)
}

func (c *StatsController) getFilter() *objs.StatsFilter {
	filter := objs.StatsFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 통계 종류
	filter.StatsType = c.Ctx.Input.Param(":statsType")

	// 자산 종류
	filter.AssetType = c.Ctx.Input.Param(":assetType")

	// 통계 조회대상 설정
	filter.OrgId, _ = strconv.Atoi(c.Ctx.Input.Param(":orgId"))
	filter.GroupId, _ = strconv.Atoi(c.Ctx.Input.Param(":groupId"))

	// Top N
	if filter.Top < 1 {
		filter.Top = 5
	}

	return &filter
}

func (c *StatsController) GetTimeline() {
	filter := c.getFilter()
	rows := c.getStatsByOrgGroup(filter, "timeline2")

	// Case 1
	timeline := make(map[string]map[int]int)
	for _, r := range rows {
		if _, ok := timeline[r.Item]; !ok {
			timeline[r.Item] = map[int]int{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
			}
		}
		timeline[r.Item][objs.StartEvent] += r.StartupCount
		timeline[r.Item][objs.ShockEvent] += r.ShockCount
		timeline[r.Item][objs.SpeedingEvent] += r.SpeedingCount
		timeline[r.Item][objs.ProximityEvent] += r.ProximityCount
	}

	//timelineByType :=
	//c.Data["json"] = timeline

	//// Case 2
	type val struct {
		name  string
		Value [2]int64 `json:"value"`
		Text  string   `json:"text"`
	}
	//timelineByType := map[string]map[string][]val{
	//	"startup":   make(map[string][]val),
	//	"shock":     make(map[string][]val),
	//	"speeding":  make(map[string][]val),
	//	"proximity": make(map[string][]val),
	//}
	timelineByType := map[string][]val{
		"startup":   make([]val, 0),
		"shock":     make([]val, 0),
		"speeding":  make([]val, 0),
		"proximity": make([]val, 0),
	}

	for date, m := range timeline {
		d := date[0:19]
		t, _ := time.Parse(ipasserver.DateDefault, d)
		timelineByType["startup"] = append(timelineByType["startup"], val{"startup", [2]int64{t.Unix() * 1000, int64(m[objs.StartEvent])}, date})
		timelineByType["shock"] = append(timelineByType["shock"], val{"shock", [2]int64{t.Unix() * 1000, int64(m[objs.ShockEvent])}, date})
		timelineByType["speeding"] = append(timelineByType["speeding"], val{"speeding", [2]int64{t.Unix() * 1000, int64(m[objs.SpeedingEvent])}, date})
		timelineByType["proximity"] = append(timelineByType["proximity"], val{"proximity", [2]int64{t.Unix() * 1000, int64(m[objs.ProximityEvent])}, date})

		//if _, ok := timelineByType["startup"][d]; !ok {
		//	timelineByType["startup"][d] = make([]val, 0)
		//	timelineByType["shock"][d] = make([]val, 0)
		//	timelineByType["speeding"][d] = make([]val, 0)
		//	timelineByType["proximity"][d] = make([]val, 0)
		//}
		////	timeLine["proximity"][r.Item] =
		//
		//timelineByType["startup"][d] = append(timelineByType["startup"][d], val{"startup", [2]int64{t.Unix() * 1000, int64(m[objs.StartEvent])}})
		//timelineByType["shock"][d] = append(timelineByType["shock"][d], val{"shock", [2]int64{t.Unix() * 1000, int64(m[objs.ShockEvent])}})
		//timelineByType["speeding"][d] = append(timelineByType["speeding"][d], val{"speeding", [2]int64{t.Unix() * 1000, int64(m[objs.SpeedingEvent])}})
		//timelineByType["proximity"][d] = append(timelineByType["proximity"][d], val{"proximity", [2]int64{t.Unix() * 1000, int64(m[objs.ProximityEvent])}})

		//	//timeLine["shock"] = append(timeLine["shock"], val{"shock", [2]int64{t.Unix() * 1000,int64(r.ShockCount)}})
		//	//timeLine["speeding"] = append(timeLine["speeding"], val{"speeding", [2]int64{t.Unix() * 1000,int64(r.SpeedingCount)}})
		//	//timeLine["proximity"] = append(timeLine["proximity"], val{"proximity", [2]int64{t.Unix() * 1000,int64(r.ProximityCount)}})
		//	//timeLine["startup"] = append(timeLine["startup"], val{t.Unix() * 1000,int64(r.StartupCount)})
		//	//timeLine["shock"] = append(timeLine["shock"], val{t.Unix() * 1000,int64(r.ShockCount)})
		//	//timeLine["speeding"] = append(timeLine["speeding"], val{t.Unix() * 1000,int64(r.SpeedingCount)})
		//	//timeLine["proximity"] = append(timeLine["proximity"], val{t.Unix() * 1000,int64(r.ProximityCount)})
	}
	//
	c.Data["json"] = timelineByType
	c.ServeJSON()
}

func (c *StatsController) GetStatsBy() {
	//http://127.0.0.1/stats/evt1/by/equip/org/-1/group/-1
	//http://127.0.0.1/stats/evt1/by/equip/org/1/group/-1
	//http://127.0.0.1/stats/evt1/by/equip/org/1/group/7

	//http://127.0.0.1/stats/evt1/by/group/org/-1/group/-1
	//http://127.0.0.1/stats/evt1/by/group/org/1/group/-1
	//http://127.0.0.1/stats/evt1/by/group/org/1/group/7

	filter := c.getFilter()
	rows, _, err := models.GetStatsBy(c.member, filter)
	if err != nil {
		log.Error(err)
	}

	if rows == nil {
		c.Data["json"] = []int{}
	} else {
		if filter.AssetType == "group" {
			c.updateItemText(rows)
		}
		c.Data["json"] = rows
	}

	c.ServeJSON()
}

func (c *StatsController) updateItemText(rows []objs.Stats) {
	for i, r := range rows {
		asset := strings.SplitN(r.Item, "/", 2)
		orgId, _ := strconv.Atoi(asset[0])
		orgAsset, ok := assetMap.Load(orgId)
		if ok {
			rows[i].OrgName = orgAsset.(objs.Asset).Name
		} else {
			rows[i].OrgName = asset[0]
		}

		groupId, _ := strconv.Atoi(asset[1])
		groupAsset, ok := assetMap.Load(groupId)
		if ok {
			rows[i].GroupName += groupAsset.(objs.Asset).Name
		} else {
			rows[i].GroupName += asset[1]
		}
	}
}

//
//func (c *StatsController) GetStats() {
//	filter := c.getFilter()
//	rows, _, err := models.GetStats(c.member, filter)
//	if err != nil {
//		log.Error(err)
//	}
//
//	if rows == nil {
//		c.Data["json"] = []int{}
//	} else {
//		c.Data["json"] = rows
//	}
//	c.ServeJSON()
//}

func (c *StatsController) GetSummary() {
	filter := c.getFilter()
	c.Data["json"] = map[string]interface{}{
		"eventTypes":       c.getEventTypes(filter),
		"equipCountByType": c.getEquipCountByType(filter),
		"activated":        c.getStatsByOrgGroup(filter, "activated"),
		"shocklinks":       c.getShockLinks(filter),
	}
	c.ServeJSON()
}

func (c *StatsController) getEventTypes(filter *objs.StatsFilter) map[int]int {
	eventTypes := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}
	filter.StatsType = "evt"
	filter.Top = 99999
	rows, _, err := models.GetStatsByAssetId(c.member, filter)
	if err != nil {
		log.Error(err)
	}
	for _, r := range rows {
		eType, _ := strconv.Atoi(r.Item)
		eventTypes[eType] += r.Count
	}
	return eventTypes
}

func (c *StatsController) getEquipCountByType(filter *objs.StatsFilter) map[int]int {
	tags := map[int]int{
		1: 0,
		2: 0,
		4: 0,
	}
	rows := c.getStatsByOrgGroup(filter, "equip_count")

	for _, r := range rows {
		equipType, _ := strconv.Atoi(r.Item)
		tags[equipType] += r.Count
	}
	return tags
}

func (c *StatsController) getStatsByOrgGroup(filter *objs.StatsFilter, statsType string) []objs.Stats {
	filter.StatsType = statsType
	rows, err := models.GetStatsByOrgGroup(c.member, filter)
	if err != nil {
		log.Error(err)
	}

	for i, r := range rows {
		rows[i].OrgName, rows[i].GroupName = GetOrgGroupName(r.OrgId, r.GroupId)
	}

	return rows
}

func (c *StatsController) getShockLinks(filter *objs.StatsFilter) []*node {
	m := make(map[string]*node)
	//nodes := make([]node, 0)
	rows := c.getStatsByOrgGroup(filter, "shocklinks")

	//arr := make([]string, 0)
	for _, r := range rows {
		code := GetOrgCode(r.OrgId)
		arr := strings.Split(r.Item, ",")
		for _, a := range arr {
			p := strings.SplitN(a, "/", 2)
			src, dst := code+"."+p[0], code+"."+p[1]

			// Left side
			if _, ok := m[src]; !ok {
				m[src] = &node{src, 1, []string{dst}}
			} else {
				m[src].Imports = append(m[src].Imports, dst)
			}

			// Right side
			if _, ok := m[dst]; !ok {
				m[dst] = &node{dst, 1, []string{src}}
			} else {
				m[dst].Imports = append(m[dst].Imports, src)
			}
		}
	}
	nodes := make([]*node, 0, len(m))
	for _, value := range m {
		nodes = append(nodes, value)
	}

	return nodes
}
