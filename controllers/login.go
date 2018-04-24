package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/ipas-mcs/libs"
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
		uri := beego.AppConfig.DefaultString("home_uri", "/ipaslogs")
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

	// 사용자 정보 조회
	member, err := models.GetMember(map[string]interface{}{
		"username": username,
	})

	if err != nil {
		log.Error(err)
		result.Message = c.Tr("msg.fail_to_request_open") + " (-3)"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	// 로그인 시도회수가 임계값 이상이이고
	// 마지막 로그인 시도 이후 x초 이내이면 접속시도 제한
	v1, _ := objs.GlobalConfig.Load("login_max_failed_login_attempts")
	maxFailedLoginAttempts := v1.(objs.MultiValue)
	v2, _ := objs.GlobalConfig.Load("login_block_seconds")
	loginBlockedTime := v2.(objs.MultiValue)
	if maxFailedLoginAttempts.ValueN > 0 && int(member.FailedLoginCount) > maxFailedLoginAttempts.ValueN { // 로그인 시도회수 초과
		elapsedTime := time.Now().Sub(member.LastFailedLogin).Seconds()
		if elapsedTime < float64(loginBlockedTime.ValueN) { // 로그인 제한시간 이내이면
			models.LoginFailed(username, false)
			result.Message = c.Tr("msg.fail_to_request_open") + " (-4)" // 로그인 시간제한
			c.audit("signin_failed", map[string]string{"username": username, "message": "locked down account"}, nil)
			c.Data["json"] = result
			c.ServeJSON()
			return
		}

	}

	// 비밀번호 비교
	userPassword, _ := hex.DecodeString(encPassword)
	serverPassword := sha256.Sum256([]byte(member.Password + member.Salt))
	if bytes.Equal(userPassword, serverPassword[:]) { // 로그인 성공
		result.State = true

		// 세정보 등록
		c.SetSession("memberId", member.MemberId)
		c.SetSession("username", member.Username)
		c.SetSession(beego.BConfig.WebConfig.Session.SessionName, c.Ctx.GetCookie(beego.BConfig.WebConfig.Session.SessionName))
		member.SessionId = c.GetSession(beego.BConfig.WebConfig.Session.SessionName).(string)
		models.AfterSignin(member)

		// 로그인 성공하면
		// 감사이력 생성
		// 로그인 실패수 초기화
		// 마지막 로그인 시간 기록

		redirectUri := c.GetString("redirectUri")
		if len(redirectUri) < 1 {
			redirectUri = beego.AppConfig.DefaultString("home_url", "/syslog")
		}
		result.Data = map[string]string{
			"redirectUrl": redirectUri,
		}
	} else {
		c.audit("signin_failed", map[string]string{"username": username, "message": "wrong password"}, nil)
		models.LoginFailed(username, true)
		result.Message = c.Tr("msg.fail_to_request_open") + " (-5)" // 비밀번호 오류
	}

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LoginController) GetPasswordSalt() {
	result := objs.NewResult()
	username := c.Ctx.Input.Param(":username")
	spew.Dump()

	// 사용자 정보 조회
	member, err := models.GetMember(map[string]interface{}{
		"username": username,
	})
	if err != nil || member.MemberId < 1 {
		if err != orm.ErrNoRows { // 예상되지 않은 에러이면 출력
			log.Error(err)
		}
		c.audit("signin_failed", map[string]string{"username": username, "message": "user not found"}, nil)
		result.Message = c.Tr("msg.fail_to_request_open") + " (-1)"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	// 비밀번호 Salt 값 생성
	salt := libs.GetRandomString(10)
	_, err = models.UpdateRow("mbr_password", "member_id", member.MemberId, map[string]interface{}{"salt": salt})
	if err != nil {
		result.Message = c.Tr("msg.fail_to_request_open") + " (-2)"
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
