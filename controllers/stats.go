package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

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
		"eventTypes":        c.getEventTypes(filter),
		"equipCountByType": c.getEquipCountByType(filter),
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
	rows, _, err := models.GetStats(c.member, filter)
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
	rows, err := models.GetEquipCountByType(c.member, filter)
	if err != nil {
		log.Error(err)
	}

	for _, r := range rows {
		tags[r.EquipType] += r.Count
	}
	return tags
}
