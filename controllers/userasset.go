package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type UserassetController struct {
	baseController
}

func (c *UserassetController) CtrlPrepare() {
	// 권한 부여
	c.grant(objs.User)
}

// 선택된 자산의 전체 하위노드 조회
func (c *UserassetController) GetChildren() {

	// 조회할 자산 그룹
	class, _ := strconv.Atoi(c.Ctx.Input.Param(":class"))

	// 자산 맵 구성
	assetMap := getUserassetMapByClassId(class, c.member)

	// 하위자산 설정
	c.Data["json"] = assetMap[RootId].Children
	c.ServeJSON()
}

// 자산 맵 조회
func getUserassetMapByClassId(class int, member *objs.Member) objs.AssetMap {
	// 클래스에 해당하는 자산 조회
	list, err := models.GetUserassetsByClass(class, member)
	CheckError(err)
	assets := make([]*objs.Asset, 0)
	for idx := range list {
		assets = append(assets, &list[idx])
	}

	// 자산 Map 구성
	assetMap := organizeAssets(class, assets)
	return assetMap
}

//
func (c *UserassetController) GetIpasList() {
	// 요청값 분류
	filter := c.getIpasFilter()
	list, total, err := models.GetIpaslist(c.member, filter)
	c.serveResultJson(list, total, err, filter.FastPaging)
}

func (c *UserassetController) getIpasFilter() *objs.IpasFilter {

	// 요청값 분류
	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	// 페이징 처리
	if filter.Sort == "" {
		filter.Sort = "equip_id"
	}
	if filter.Order == "" {
		filter.Order = "asc"
	}
	if filter.Limit < 1 {
		filter.Limit = 20
	}
	if filter.FastPaging == "" {
		filter.FastPaging = "on"
	}
	return &filter
}


