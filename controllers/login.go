package controllers

//import (
//	log "github.com/sirupsen/logrus"
//)

type LoginController struct {
	baseController
}

func (c *LoginController) SubPrepare() {
	c.isLoginRequired = false
}

func (c *LoginController) Get() {
	c.setTpl("login.tpl")
}

func (c *LoginController) GetPasswordSalt() {
	//username := c.Ctx.Input.Param(":username")

	//// Check if member exists
	//member, err := models.GetMemberByUsername(username)
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
	c.ServeJSON()
}