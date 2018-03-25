package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/devplayg/ipas-mcs/libs"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"strings"
)

type MemberController struct {
	baseController
}

func (c *MemberController) LangPrepare() {

	c.langMap["cancel"] = c.Tr("cancel")
}

// 사용자 정보 출력
func (c *MemberController) Get() {
	if c.IsAjax() { // Ajax 요청이면 Json 타입으로 리턴
		if c.Ctx.Input.Param(":memberId") != "" { // 특정 사용자 정보 요청이면
			c.GetMemberById()
		} else {
			filter := c.getPagingFilter()
			members, total, err := models.GetMembers(filter)
			c.serveResultJson(members, total, err, "off")
		}

	} else { // Ajax 외 요청이면 HTML 리턴
		positions := make(map[string]int)

		positions["Administrator"] = objs.Administrator
		positions["Superman"] = objs.Superman
		positions["Observer"] = objs.Observer
		c.Data["positions"] = positions

		//c.langToFrontEnd("msg.confirm_delete")
		c.setTpl("member.tpl")
		//c.Data["langMap"].(map[string]string)["cancel"] = "abc"

	}
}

// 사용자 정보 조회
func (c *MemberController) GetMemberById() {
	dbResult := objs.NewDbResult()
	c.Data["json"] = dbResult

	// 사용자 정보 조회
	memberId, _ := strconv.Atoi(c.Ctx.Input.Param(":memberId"))
	m, err := models.GetMember(map[string]interface{}{
		"t.member_id": memberId,
	})
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
		dbResult.Data = m
	}
	c.ServeJSON()
}

// 사용자 등록
func (c *MemberController) Post() {
	dbResult := objs.NewDbResult()

	member := objs.Member{}
	if err := c.ParseForm(&member); err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}
	member.Username = strings.ToLower(member.Username) // 아이디는 소문자로
	member.Position |= objs.User                       // "일반"권한은 기본 추가

	if err := c.CheckForm(&member); err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}

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

// 입력폼 체크
func (c *MemberController) CheckForm(m *objs.Member) error {
	for _, g := range m.UserGroups {
		if 1<<uint(g) > objs.Administrator {
			return errors.New("unauthorized user group")
		}
	}

	// 권한 설정
	for _, v := range m.UserGroups {
		if v > 0 && v < objs.Superman {
			m.Position |= 1 << uint(v)
		}
	}

	// 비밀번호 설정
	encPassword := sha256.Sum256([]byte(m.Username + m.Password))
	m.PasswordConfirm = hex.EncodeToString(encPassword[:])

	// IP ACL 설정
	m.AllowedIpList = make([]objs.IpCidr, 0)
	list := libs.SplitString(m.AllowedIp, `[\s|,]+`)
	for _, s := range list {
		if strings.Index(s, "/") > -1 { // CIDR이 있으면
			ip, ipNet, err := net.ParseCIDR(s)
			if err != nil {
				return err
			}
			mask, _ := ipNet.Mask.Size()
			m.AllowedIpList = append(m.AllowedIpList, objs.IpCidr{ip.String(), mask})
		} else { // IP면
			ip := net.ParseIP(s)
			if ip == nil {
				return errors.New("invalid IP address: " + s)
			}
			m.AllowedIpList = append(m.AllowedIpList, objs.IpCidr{ip.String(), 32})
		}
	}

	// 입력값 유효성 체크
	if err := m.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *MemberController) getPagingFilter() *objs.PagingFilter {

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

// 사용자 정보 업데이트
func (c *MemberController) Patch() {
	dbResult := objs.NewDbResult()

	member := objs.Member{}
	if err := c.ParseForm(&member); err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}
	if err := c.CheckForm(&member); err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}

	_, err := models.UpdateMember(&member, c.member)
	if err != nil {
		dbResult.Message = err.Error()
	} else {
		dbResult.State = true
	}
	c.Data["json"] = dbResult
	c.ServeJSON()
}

func (c *MemberController) Delete() {
	dbResult := objs.NewDbResult()

	memberId, _ := strconv.Atoi(c.Ctx.Input.Param(":memberId"))
	rs, err := models.RemoveMember(memberId, c.member.Position)
	if err != nil {
		dbResult.Message = err.Error()
		c.Data["json"] = dbResult
		c.ServeJSON()
		return
	}

	affectedRows, _ := rs.RowsAffected()
	if affectedRows > 0 {
		dbResult.State = true
	} else {
		dbResult.Message = c.Tr("msg.not_founded")
		dbResult.Message += " / " + c.Tr("msg.not_permitted")
	}

	c.Data["json"] = dbResult
	c.ServeJSON()
}


func (c *MemberController) GetMemberAcl() {
	//dbResult := objs.NewDbResult()
	//
	//memberId, _ := strconv.Atoi(c.Ctx.Input.Param(":memberId"))
	//
	//assets, err := models.GetMemberAcl(memberId)
	//c.Data["json"]  = assets
	//c.jsons
}