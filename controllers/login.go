package controllers

import (
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/sirupsen/logrus"
)

type LoginController struct {
	baseController
}

func (c *LoginController) CtrlPrepare() {
	c.loginRequired(false)
}

func (c *LoginController) Get() {
	c.setTpl("login.tpl")
}

func (c *LoginController) GetPasswordSalt() {
	result := objs.NewResult()

	username := c.Ctx.Input.Param(":username")

	// Check if member exists
	member, err := models.GetMemberByUsername(username)
	if err != nil {
		result.Message = err.Error()
		c.ServeJSON()
		return
	}

	if member != nil {
		salt := GetRandomString(10)
		result.Data = salt
		logrus.Debugf("Salt: %s", salt)
		result.State = true

		data := map[string]interface{} {
			"zipcode": salt,
			"country": salt+salt,
			"state": len(salt),
			"last_read_message":len(salt),
		}

		//rs, err := models.UpdateMember(member.MemberId, data )
		rs, err := models.UpdateRow("mbr_member", "member_id", member.MemberId, data)
		checkErr(err)
		affectedRows, _ := rs.RowsAffected()
		logrus.Debugf("AffectedRows: %d", affectedRows)

	} else {
		result.Message = "No users"
	}
	//if member.MemberId < 1 || err != nil {
	//	result := models.Result{false, c.Tr("msg_fail_to_request_open"), "1"}
	//	c.Data["json"] = result
	//	c.WriteAuditLog("signin_failed", result, "")
	//} else {
	//	salt := libs.GetRandomString(10)
	//	_, err := models.UpdatePassword(member.MemberId, "", salt)
	//	CheckError(err)
	//	c.Data["json"] = models.Result{true, "", salt}
	//}
	c.Data["json"] = result
	c.ServeJSON()
}
