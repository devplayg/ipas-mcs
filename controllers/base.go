package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"html/template"
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
	beego.Controller               // 메인 구조체 임베딩
	i18n.Locale                    // 다국어
	isLoginRequired  bool          // 로그인 필수 여부
	acl              int           // 권한
	//member           models.Member // 사용자 정보
	isLogged         bool          // 로그인 상태
	ctrlName         string        // Controller 이름
	actName          string        // Action 이름
}

func (c *baseController) Prepare() {

	// 기본 접근권한 설정
	c.isLoginRequired = true         // 로그인 필수
	c.grant(Superman, Administrator) // 관리자 이상만 실행 허용

	// 호출된 Controller 접근권한 덮어쓰기
	if app, ok := c.AppController.(CtrlPreparer); ok {
		app.CtrlPrepare()
	}

	// Client 상태 업데이트

	// 접근권한 검토

	// Controller 와 Action 이름 설정
	c.ctrlName, c.actName = c.GetControllerAndAction()
	c.ctrlName = strings.ToLower(strings.TrimSuffix(c.ctrlName, "Controller"))

	// 언어 설정
	c.setLangVer()

	// 기본 템플릿 변수 설정
	c.Data["title"] = beego.BConfig.AppName
	//c.Data["member"] = c.GetMember()
	//c.Data["IsLogged"] = c.IsLogged
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["ctrl"] = c.ctrlName
	c.Data["act"] = c.actName
}

func (c *baseController) setTpl(tplName string) {
	c.TplName = c.ctrlName + "/" + tplName
}

func (c *baseController) grant(auth ...int) {
	for _, n := range auth {
		c.acl |= n
	}
}

func (c *baseController) setLangVer() {
	hasCookie := false

	// Check if cookies have language settings
	lang := c.Ctx.GetCookie("lang")
	if len(lang) > 0 {
		if i18n.IsExist(lang) {
			hasCookie = true
		}
	}

	//  Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	if len(lang) == 0 {
		lang = "en-us"
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		c.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	c.Lang = lang
	c.Data["Lang"] = curLang.Lang
	c.Data["CurLang"] = curLang.Name
	c.Data["RestLangs"] = restLangs
}
