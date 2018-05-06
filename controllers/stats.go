package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"github.com/devplayg/ipas-mcs/models"
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

func (c *StatsController) GetStats() {
	rows, _, err := c.getStats()
	if err != nil || len(rows) == 0 { // 에러가 발생했거나, 데이터가 없으면
		log.Error(err)
		c.Data["json"] = make([]int, 0) // 크기가 0인 배열 출력
	} else { // 데이터가 있으면
		if c.Ctx.Input.Param(":assetType") == "group" {
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
		}
		c.Data["json"] = rows
	}

	c.ServeJSON()
}

func (c *StatsController) getStats() ([]objs.Stats, int64, error) {
	t := time.Now()
	filter := make(map[string]interface{})

	// 라우팅 값 설정
	filter["statsType"] = c.Ctx.Input.Param(":statsType")
	filter["assetType"] = c.Ctx.Input.Param(":assetType")

	// 자산 ID 설정
	assetId, _ := c.GetInt("assetId")
	if assetId == 0 { // 입력값이 없으면
		if c.member.Position < objs.Administrator { // 관리자가 아니면
			filter["assetId"] = c.member.MemberId * -1 // 사용자 통계값 참조
		} else {
			filter["assetId"] = -1 // 관리자 통계값 참조
		}
	} else {
		filter["assetId"] = assetId
	}

	// 날짜 설정
	if len(c.GetString("from")) > 0 || len(c.GetString("to")) > 0 {
		filter["from"] = c.GetString("from") + " 00:00:00"
		filter["to"] = c.GetString("to") + " 23:59:59"
	} else {
		filter["from"] = t.Format("2006-01-02") + " 00:00:00"
		filter["to"] = t.Format("2006-01-02") + " 23:59:59"
	}

	// Top
	top, err := c.GetInt("top")
	if err != nil || top < 1 {
		filter["top"] = 3
	} else {
		filter["top"] = top
	}

	rows, total, err := models.GetStats(filter)
	return rows, total, err
}
