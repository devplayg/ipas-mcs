package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-mcs/models"
	"strconv"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"strings"
)

type IpaslistController struct {
	baseController
}

func (c *IpaslistController) GetIpasInOrg() {
	filter := c.getFilter()
	spew.Dump(filter)
	logs, total, err := models.GetIpaslist(filter)
	//if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
	//	//filter := c.getFilter()
	//	//logs, total, err := models.GetSamplelog(filter)
	c.serveResultJson(logs, total, err, filter.FastPaging)
	//} else { // Ajax 외 요청이면 HTML 리턴
	//	//c.Data["filter"] = filter
	//	c.setTpl("asset.tpl")
	//}
}

func (c *IpaslistController) GetIpasInGroup() {
	filter := c.getFilter()
	spew.Dump(filter)
	logs, total, err := models.GetIpaslist(filter)
	//if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
	//	//filter := c.getFilter()
	//	//logs, total, err := models.GetSamplelog(filter)
	c.serveResultJson(logs, total, err, filter.FastPaging)
	//} else { // Ajax 외 요청이면 HTML 리턴
	//	//c.Data["filter"] = filter
	//	c.setTpl("asset.tpl")
	//}
}

func (c *IpaslistController) getFilter() *objs.IpasFilter {

	filter := objs.IpasFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	orgId, _ := strconv.Atoi(c.Ctx.Input.Param(":orgId"))
	if orgId > 0 {
		filter.OrgId = append(filter.OrgId, orgId)
	}

	groupId, _ := strconv.Atoi(c.Ctx.Input.Param(":groupId"))
	if groupId > 0 {
		filter.GroupId = append(filter.GroupId, groupId)
	}
	spew.Dump(filter)
	//logrus.Debug(assetId)

	// 페이징 처리
	if filter.Sort == "" {
		filter.Sort = "equip_id"
	}
	if filter.Order == "" {
		filter.Order = "asc"
	}
	if filter.Limit < 1 {
		filter.Limit = 15
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}

func (c *IpaslistController) UpdateIpasGroup() {

	groupId, _ := strconv.Atoi(c.Ctx.Input.Param(":groupId"))

	type input struct {
		List []string `form:"list[]"`
	}

	form := input{}
	if err := c.ParseForm(&form); err != nil {
		CheckError(err)
	}

	list := make([]objs.Ipas, 0)
	for _, s := range form.List {
		arr := strings.SplitN(s, "/", 2)
		orgId, _ := strconv.Atoi(arr[0])
		list = append(list, objs.Ipas{
			OrgId: orgId,
			EquipId: arr[1],
		})
	}
	//spew.Dump(list)

	dbResult := objs.NewDbResult()
	rs, err := models.UpdateIpasGroup(groupId, list)
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
		dbResult.AffectedRows, _ = rs.RowsAffected()
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}