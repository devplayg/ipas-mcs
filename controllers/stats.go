package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"strconv"
	"strings"
	"time"
)

type StatsController struct {
	baseController
}

func (c *StatsController) CtrlPrepare() {
	// 권한 부여
	c.grant(objs.User)
}

func (c *StatsController) getFilter() map[string]interface{} {
	t := time.Now()
	filter := make(map[string]interface{})

	// 라우팅 값 설정
	filter["statsType"] = c.Ctx.Input.Param(":statsType")
	filter["assetType"] = c.Ctx.Input.Param(":assetType")

	// 날짜 설정
	if len(c.GetString("from")) > 0 || len(c.GetString("to")) > 0 {
		filter["from"] = c.GetString("from") + " 00:00:00"
		filter["to"] = c.GetString("to") + " 23:59:59"
	} else {
		filter["from"] = t.Format("2006-01-02") + " 00:00:00"
		filter["to"] = t.Format("2006-01-02") + " 23:59:59"
	}

	// "그룹"통계에서만 사용
	// 자산 키 (기관코드/그룹코드)
	// 0/-1 : 전체
	// 1/0: 기관코드가 1이고, 그룹이 미분류인 통계
	// 1/2: 기관코드가 1이고, 그룹코드가 2인 통계
	filter["assetKey"] = c.GetString("assetKey")
	asset := strings.SplitN(c.GetString("assetKey"), "/", 2)
	if len(asset) == 2 {
		filter["orgId"] = asset[0]
		filter["groupId"] = asset[1]
	} else {
		// 전체 선택
		filter["orgId"] = "0"
		filter["groupId"] = "-1"
	}

	// 통계 자산 선택
	filter["assetId"],_ = c.GetInt("assetId")

	// Top
	top, err := c.GetInt("top")
	if err != nil || top < 1 {
		filter["top"] = 3
	} else {
		filter["top"] = top
	}

	return filter
}

func (c *StatsController) GetOrgGroupStats() {
	filter := c.getFilter()
	
	// 통계 조회
	rows, _, err := models.GetOrgGroupStats(c.member, filter)
	if err != nil || len(rows) == 0 { // 에러가 발생했거나, 데이터가 없으면
		c.Data["json"] = make([]int, 0) // 크기가 0인 배열 출력
	} else { // 데이터가 있으면
		for i, r := range rows {
			asset := strings.SplitN(r.Item, "/", 2)
			orgId, _ := strconv.Atoi(asset[0])
			orgAsset, ok := assetMap.Load(orgId)
			if ok {
				rows[i].ItemText = orgAsset.(objs.Asset).Name
			} else {
				rows[i].ItemText = asset[0]
			}

			groupId, _ := strconv.Atoi(asset[1])
			groupAsset, ok := assetMap.Load(groupId)
			if ok {
				rows[i].ItemText += " / " + groupAsset.(objs.Asset).Name
			} else {
				rows[i].ItemText += " / " + asset[1]
			}
		}
		c.Data["json"] = rows
	}

	c.ServeJSON()
}

func (c *StatsController) GetEquipStats() {
	filter := c.getFilter()

	// 통계 조회
	rows, _, err := models.GetStats(c.member, filter)
	if err != nil || len(rows) == 0 { // 에러가 발생했거나, 데이터가 없으면
		c.Data["json"] = make([]int, 0) // 크기가 0인 배열 출력
	} else { // 데이터가 있으면
		c.Data["json"] = rows
	}

	c.ServeJSON()
}
