package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"strings"
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
func (c *MemberController) Post() {
	member := objs.Member{}
	if err := c.ParseForm(&member); err != nil {
		log.Error(err)
	}

	member.Username = strings.ToLower(member.Username) // 아이디는 소문자로
	member.Position = objs.User                        // 권한을 "일반"으로 등록
	encPassword := sha256.Sum256([]byte(member.Username + member.Password))
	member.PasswordConfirm = hex.EncodeToString(encPassword[:])

	dbResult := objs.NewDbResult()
	rs, err := models.AddMember(&member)
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.AffectedRows, _ = rs.RowsAffected()
		if dbResult.AffectedRows == 1 {
			dbResult.State = true
		}
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

func (c *MemberController) getFilter() *objs.PagingFilter {

	// 요청값 분류
	filter := objs.PagingFilter{}
	if err := c.ParseForm(&filter); err != nil {
		log.Error(err)
	}

	if filter.FastPaging == "" {
		filter.FastPaging = "off"
	}

	return &filter
}
