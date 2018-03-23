package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"strconv"
	"github.com/astaxie/beego/logs"
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

//
func (c *AssetController) GetDescendants() {

	// 가져올 자산 class
	class, _ := strconv.Atoi(c.Ctx.Input.Param(":class"))

	// 최상위 자산 ID
	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))

	assetMap := getAssetMapByClassId(class)

	// Print
	if assetMap[assetId].Children == nil {
		c.Data["json"] = []int{}
	} else {
		c.Data["json"] = assetMap[assetId].Children
	}
	c.ServeJSON()
}


func (c *AssetController) GetChildren() {

	// Get information
	assetId, _ := strconv.Atoi(c.Ctx.Input.Param(":assetId"))

	// Get assets
	assets, err := models.GetAssetChildren(assetId)
	logs.Error(err)

	// Print
	c.Data["json"] = assets
	c.ServeJSON()
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
func organizeAssets(assets []*objs.Asset) map[int]*objs.Asset {

	// Create a root node
	assetMap := make(map[int]*objs.Asset)
	root := objs.Asset{
		AssetId: 0,
		Id: "assetid_0",
		Class: 1,
		ParentId: -1,
		Type: "type_0",
		Type1: 0,
		Type2: 0,
		Text: "root",
		Children: nil,
	}
	assetMap[0] = &root

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
func getAssetMapByClassId(class int) map[int]*objs.Asset {
	var assetMap map[int]*objs.Asset

	list, err := models.GetAssetsByClass(class)
	CheckError(err)
	assets := make([]*objs.Asset, 0)
	for idx, _ := range list {
		assets = append(assets, &list[idx])
	}
		assetMap = organizeAssets(assets)
	return assetMap
}

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
