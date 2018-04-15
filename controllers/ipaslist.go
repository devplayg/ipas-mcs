package controllers

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/devplayg/ipas-mcs/models"
	"strconv"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
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