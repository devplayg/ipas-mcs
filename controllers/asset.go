package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"strconv"
)

const (
	RootId = 0
)

type AssetController struct {
	baseController
}

func (c *AssetController) Get() {
	//filter := c.getFilter()

	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		//filter := c.getFilter()
		//logs, total, err := models.GetSamplelog(filter)
		//c.serveResultJson(logs, total, err, filter.FastPaging)
	} else { // Ajax 외 요청이면 HTML 리턴
		//c.Data["filter"] = filter
		c.setTpl("asset.tpl")
	}
}

func (c *AssetController) Post() {

	// Parse form
	asset := objs.Asset{}
	if err := c.ParseForm(&asset); err != nil {
		result := objs.NewResult()
		result.Message = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	_, err := models.AddAsset(asset)
	dbResult := objs.NewDbResult()
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

// 선택된 자산의 전체 하위노드 조회
//func (c *AssetController) GetChildren() {
//
//	// Get information
//	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))
//
//	// Get assets
//	assets, err := models.GetAssetChildren(assetId)
//	logs.Error(err)
//
//	// Print
//	c.Data["json"] = assets
//	c.ServeJSON()
//}

// 선택된 자산의 전체 하위노드 조회
func (c *AssetController) GetDescendants() {

	// 조회할 자산 그룹
	class, _ := strconv.Atoi(c.Ctx.Input.Param(":class"))

	// 자산 ID
	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))

	// 자산 맵 구성
	asset := c.getDescendants(class, assetId)

	if asset != nil {
		c.Data["json"] = asset.Children
	}
	c.ServeJSON()
}

func (c *AssetController) GetDescendantsWithRoot() {

	// 조회할 자산 그룹
	class, _ := strconv.Atoi(c.Ctx.Input.Param(":class"))

	// 자산 ID
	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))

	// 자산 맵 구성
	assetMap := getAssetMapByClassId(class)

	if assetId == RootId {
		c.Data["json"] = assetMap[assetId]
	} else {
		root := objs.NewRootAsset(class)
		root.Children = append(root.Children, assetMap[assetId])
		c.Data["json"] = root
	}
	c.ServeJSON()
}

func (c *AssetController) getDescendants(class, assetId int) *objs.Asset {
	assetMap := getAssetMapByClassId(class)
	return assetMap[assetId]
}

// 자산ID로 조회
func (c *AssetController) GetAsset() {
	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))
	asset, err := models.GetAsset(assetId)
	dbResult := objs.NewDbResult()
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
		dbResult.Data = asset
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

// 자산정보 업데이트
func (c *AssetController) UpdateAsset() {
	asset := objs.Asset{}
	if err := c.ParseForm(&asset); err != nil {
		CheckError(err)
	}
	rs, err := models.UpdateAsset(&asset)

	dbResult := objs.NewDbResult()
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
		dbResult.AffectedRows, _ = rs.RowsAffected()
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

// 자산정보 삭제
func (c *AssetController) RemoveAsset() {
	type input struct {
		AssetIdList []int `form:"asset_id_list[]"`
	}

	target := input{}
	if err := c.ParseForm(&target); err != nil {
		CheckError(err)
	}

	dbResult := objs.NewDbResult()
	rs, err := models.RemoveAsset(target.AssetIdList)
	if err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}

	var affectedRows int64
	if rs != nil {
		affectedRows, _ = rs.RowsAffected()
	}
	if affectedRows > 0 {
		dbResult.State = true
		dbResult.AffectedRows = affectedRows
	} else {
		dbResult.Message = c.Tr("msg.not_founded")
	}

	c.Data["json"] = dbResult
	c.ServeJSON()
}

// 자산 맵 조회
func getAssetMapByClassId(class int) objs.AssetMap {

	// 클래스에 해당하는 자산 조회
	list, err := models.GetAssetsByClass(class)
	CheckError(err)
	assets := make([]*objs.Asset, 0)
	for idx := range list {
		assets = append(assets, &list[idx])
	}

	// 자산 Map 구성
	assetMap := organizeAssets(class, assets)
	return assetMap
}

// 자산정보 조직화
func organizeAssets(class int, assets []*objs.Asset) objs.AssetMap {

	// Create map and root node
	assetMap := make(objs.AssetMap)
	assetMap[RootId] = objs.NewRootAsset(class)
	assetMap[RootId].State.Opened = true

	var keys []int
	for idx, asset := range assets {
		assetMap[asset.AssetId] = assets[idx]
		keys = append(keys, asset.AssetId)
	}

	// Organize assets
	for _, k := range keys {
		parentId := assetMap[k].ParentId
		if _, ok := assetMap[parentId]; ok {
			assetMap[parentId].Children = append(assetMap[parentId].Children, assetMap[k])
		} else { // Lost child
			if assetMap[k].AssetId > 0 { // lost child
				parentId = 0
				assetMap[k].ParentId = parentId
				assetMap[parentId].Children = append(assetMap[parentId].Children, assetMap[k])
			}
		}
	}

	return assetMap
}

//
//func (this *AssetsController) Patch() {
//
//	// Parse form
//	asset := models.Asset{}
//	if err := this.ParseForm(&asset); err != nil {
//		result := models.Result{false, err.Error(), ""}
//		this.Data["json"] = result
//		this.ServeJSON()
//		return
//	}
//
//	// Update
//	asset.AssetId, _ = strconv.Atoi(this.Ctx.Input.Param(":assetId"))
//	_, err := models.UpdateAsset(asset)
//	logs.Error(err)
//	if err == nil {
//		this.Data["json"] = models.Result{State: true}
//	} else {
//		this.Data["json"] = models.Result{State: false, Message: err.Error()}
//	}
//	this.ServeJSON()
//}
//
//func (this *AssetsController) Delete() {
//
//	// Parse form
//	asset := models.Asset{}
//	if err := this.ParseForm(&asset); err != nil {
//		result := models.Result{false, err.Error(), ""}
//		this.Data["json"] = result
//		this.ServeJSON()
//		return
//	}
//
//	// Get assets
//	list, err := models.GetAssetsByClass(0)
//	logs.Error(err)
//	assets := make([]*models.Asset, 0)
//	for idx, _ := range list {
//		assets = append(assets, &list[idx])
//	}
//
//	// Organize assets
//	assetMap := organizeAssets(assets)
//
//	// Get assets to be deleted
//	asset.AssetId, _ = strconv.Atoi(this.Ctx.Input.Param(":assetId"))
//	list_to_be_deleted := GetAssetDescendantsIdList(assetMap, asset.AssetId)
//
//	// Delete assets
//	err = models.DeleteAsset(list_to_be_deleted)
//	logs.Error(err)
//
//	// Print
//	if err == nil {
//		this.Data["json"] = models.Result{State: true}
//	} else {
//		this.Data["json"] = models.Result{State: false, Message: err.Error()}
//	}
//	this.ServeJSON()
//	spew.Dump()
//}
//
//func (this *AssetsController) GetMyAssets() {
//
//	// Get parameters
//	class, _ := strconv.Atoi(this.Ctx.Input.Param(":class"))
//
//	// Get member's assets
//	assetMap := getMemberAssetMapByClassId(this.GetMember(), class)
//
//	// Print
//	this.Data["json"] = assetMap[0].Children
//	this.ServeJSON()
//}
//
//func getMyAssetList(memberId, class int) []*models.Asset {
//
//	// Get all assets by class (ordered)
//	assets, err := models.GetAssetsByClass(class)
//	logs.Error(err)
//
//	idMap := make(map[int]int)              // Map: asset_id/parent_id
//	assetMap := make(map[int]*models.Asset) // Map: asset_id/asset
//	for idx, asset := range assets {
//		idMap[asset.AssetId] = asset.ParentId
//		assetMap[asset.AssetId] = &assets[idx]
//	}
//
//	// Get member's assets(only leaf nodes) and track to the root node
//	memberAssetIdList, err := models.GetAssetIdListByMemberId(memberId, class)
//	logs.Error(err)
//	myAssetIdMap := make(map[int]bool, 0)
//	for _, assetId := range memberAssetIdList {
//		myAssetIdMap[assetId] = true
//
//		parentId, ok := idMap[assetId]
//		for ok && parentId != 0 {
//			myAssetIdMap[parentId] = true
//			parentId, ok = idMap[parentId]
//		}
//	}
//
//	myAssetList := make([]*models.Asset, 0)
//
//	// NOT ORDERED
//	//	for assetId, _ := range myAssetIdMap {
//	//		if _, ok := assetMap[assetId]; ok {
//	//			myAssetList = append(myAssetList, assetMap[assetId])
//	//		}
//	//	}
//
//	// ORDERED
//	for _, asset := range assets { // for sorting
//		if _, ok := myAssetIdMap[asset.AssetId]; ok {
//			myAssetList = append(myAssetList, assetMap[asset.AssetId])
//		}
//	}
//
//	return myAssetList
//}
//

//
//func getMemberAssetMapByClassId(member *models.Member, class int) map[int]*models.Asset {
//	var assetMap map[int]*models.Asset
//
//	if member.Position&Administrator > 0 || member.Position&OBSERVER > 0 { // is admin or observer
//		list, err := models.GetAssetsByClass(class)
//		CheckError(err)
//		assets := make([]*models.Asset, 0)
//		for idx, _ := range list {
//			assets = append(assets, &list[idx])
//		}
//		assetMap = organizeAssets(assets)
//	} else {
//		assets := getMyAssetList(member.MemberId, class)
//		assetMap = organizeAssets(assets)
//	}
//
//	return assetMap
//}
//

//
//func GetAssetDescendantsIdList(assetMap map[int]*models.Asset, assetId int) (id_list []int) {
//	id_list = append(id_list, assetId)
//	for _, child := range assetMap[assetId].Children {
//		list := GetAssetDescendantsIdList(assetMap, child.AssetId)
//		id_list = append(id_list, list...)
//	}
//
//	return id_list
//}
