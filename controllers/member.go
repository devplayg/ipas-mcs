package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
)

type MemberController struct {
	baseController
}

func (c *MemberController) Get() {
	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		filter := c.getFilter()
		members, total, err := models.GetMembers(filter)
		checkErr(err)
		c.serveResultJson(members, total, err, "off")
	} else { // Ajax 외 요청이면 HTML 리턴
		positions := make(map[string]int)
		positions["User"] = objs.User
		positions["Administrator"] = objs.Administrator
		positions["Superman"] = objs.Superman
		positions["Observer"] = objs.Observer
		c.Data["positions"] = positions

		c.setTpl("member.tpl")
	}
}

func (c *MemberController) getFilter() *objs.CommonFilter {

	// 요청값 분류
	filter := objs.CommonFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}




