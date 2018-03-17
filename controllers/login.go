package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"time"
)

type LoginController struct {
	baseController
}

func (c *LoginController) CtrlPrepare() {
	c.loginRequired(false)
}

// 로그인 페이지
func (c *LoginController) Get() {
	if c.isLogged {
		uri := beego.AppConfig.DefaultString("home_uri", "/ipaslog")
		c.Redirect(uri, 302)
	} else {
		if len(c.GetString("redirectUri")) > 0 {
			c.Data["redirectUri"] = c.GetString("redirectUri")
		}
	}

	c.setTpl("login.tpl")
}

// 로그인 시도
func (c *LoginController) Post() {

	// 리턴 구조체 설정
	result := objs.NewResult()

	// 입력값 수신
	username := c.GetString("username")
	encPassword := c.GetString("encPassword")
	log.Debugf("username=%s, pwd=%s", username, encPassword)

	// 사용자 정보 조회
	member, err := models.GetMember(map[string]interface{}{
		"username": username,
	})

	if err != nil { // 존재하지 않으면
		result.Message = err.Error()
	} else {
		userPassword, _ := hex.DecodeString(encPassword)
		serverPassword := sha256.Sum256([]byte(member.Password + member.Salt))

		if bytes.Equal(userPassword, serverPassword[:]) {
			result.State = true

			c.SetSession("memberId", member.MemberId)
			c.SetSession("username", member.Username)
			c.SetSession("sessionId", c.Ctx.GetCookie(beego.BConfig.WebConfig.Session.SessionName))

			models.UpdateRow(
				"mbr_member",
				"member_id",
				member.MemberId,
				map[string]interface{}{
					"last_success_login": time.Now().Format(objs.DefaultDateFormat),
					"failed_login_count": 0,
				},
			)

			// 로그인 성공하면
			// 감사이력 생성
			// 로그인 실패수 초기화
			// 마지막 로그인 시간 기록
			result.Data = map[string]string{
				"redirectUrl": beego.AppConfig.DefaultString("home_url", "/syslog"),
			}
		} else {
			result.Message = c.Tr("msg_fail_to_request_open") + " (-1)"
		}
	}

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LoginController) GetPasswordSalt() {
	result := objs.NewResult()
	username := c.Ctx.Input.Param(":username")

	// Check if member exists
	//member, err := models.GetMemberByUsername(username)
	member, err := models.GetMember(map[string]interface{}{
		"username": username,
	})
	c.audit("test_logging", map[string]string{"username": username}, nil)
	if err != nil {
		// 결과 값이 없는 상황 이외에는 시스템 로깅
		if err != orm.ErrNoRows {
			checkErr(err)
		}
		c.audit("signin_failed", map[string]string{"username": username, "message": err.Error()}, nil)
		result.Message = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	// 비밀번호 Salt 값 생성
	salt := GetRandomString(10)
	_, err = models.UpdateRow("mbr_password", "member_id", member.MemberId, map[string]interface{}{"salt": salt})
	if err != nil {
		checkErr(err)
		result.Message = err.Error()
	} else {
		result.Data = salt
		result.State = true
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LoginController) Logout() {
	c.DestroySession()
	c.Redirect("/signin", 302)
}
