package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

const (
	User = 1 << iota
	UnknownLeve2
	UnknownLeve3
	UnknownLeve4
	UnknownLeve5
	UnknownLeve6
	UnknownLeve7
	UnknownLeve8
	UnknownLeve9
	Administrator // 512
	Superman      // 1024
)

type CtrlPreparer interface {
	CtrlPrepare()
}

// 기본 Controller
type baseController struct {
	beego.Controller      // 메인 구조체 임베딩
	i18n.Locale           // 다국어
	isLoginRequired  bool // 로그인 필수 여부
	acl              int  // 권한
	//member           models.Member // 사용자 정보
	isLogged bool   // 로그인 상태
	ctrlName string // Controller 이름
	actName  string // Action 이름
}

func (c *baseController) Prepare() {
	// 기본권한 설정
	c.isLoginRequired = true         // 로그인 필수
	c.grant(Superman, Administrator) // 관리자 이상만 실행 허용

	// Controller 와 Action 이름 설정
	c.ctrlName, c.actName = c.GetControllerAndAction()
	c.ctrlName = strings.ToLower(strings.TrimSuffix(c.ctrlName, "Controller"))
}

func (c *baseController) setTpl(tplName string) {
	c.TplName = c.ctrlName + "/" + tplName
}

func (c *baseController) grant(auth ...int) {
	for _, n := range auth {
		c.acl |= n
	}
}
