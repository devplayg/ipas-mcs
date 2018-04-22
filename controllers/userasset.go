package controllers
//
//import (
//	"strconv"
//	"github.com/devplayg/ipas-mcs/objs"
//	"github.com/devplayg/ipas-mcs/models"
//)
//
//
//type UserassetController struct {
//	baseController
//}
//
//
//func (c *UserassetController) CtrlPrepare() {
//	// 권한 부여
//	c.grant(objs.User)
//}
//
//
//
//// 선택된 자산의 전체 하위노드 조회
//func (c *UserassetController) GetDescendants() {
//
//	// 조회할 자산 그룹
//	class, _ := strconv.Atoi(c.Ctx.Input.Param(":class"))
//
//	// 자산 ID
//	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))
//
//	// 자산 맵 구성
//	assetMap := c.getUserassetMapByClassId(class, assetId)
//
//	if assetMap[assetId] != nil {
//		c.Data["json"] = assetMap[assetId].Children
//	}
//	c.ServeJSON()
//}
//
//
//// 자산 맵 조회
//func (c *UserassetController)  getUserassetMapByClassId(class, assetId int) objs.AssetMap {
//
//	// 클래스에 해당하는 자산 조회
//	list, err := models.GetUserassetsByClass(class, c.member)
//	CheckError(err)
//	assets := make([]*objs.Asset, 0)
//	for idx := range list {
//		assets = append(assets, &list[idx])
//	}
//
//	// 자산 Map 구성
//	assetMap := organizeAssets(class, assets)
//	return assetMap
//}
