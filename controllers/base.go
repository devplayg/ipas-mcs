package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
	"html/template"
	"net/url"
	"github.com/davecgh/go-spew/spew"
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
	beego.Controller              // 메인 구조체 임베딩
	i18n.Locale                   // 다국어
	isLoginRequired  bool         // 로그인 필수 여부
	acl              int          // 권한
	member           *objs.Member // 사용자 정보
	isLogged         bool         // 로그인 상태
	ctrlName         string       // Controller 이름
	actName          string       // Action 이름
}

func (c *baseController) Prepare() {

	// Controller 와 Action 이름 설정
	c.ctrlName, c.actName = c.GetControllerAndAction()
	c.ctrlName = strings.ToLower(strings.TrimSuffix(c.ctrlName, "Controller"))

	// 기본 접근권한 설정
	c.isLoginRequired = true         // 로그인 필수
	c.grant(Superman, Administrator) // 관리자 이상만 실행 허용
	c.isLogged = false               // 로그인 상태

	// 호출된 Controller 접근권한 덮어쓰기
	if app, ok := c.AppController.(CtrlPreparer); ok {
		app.CtrlPrepare()
	}
	// 로그인 상태 체크
	c.checkLoginStatus()

	// 요청 디버깅 코드
	if beego.BConfig.RunMode == "dev" {
		spew.Println("=============START=========================================================")
		log.Debugf("Method=%s, Ctrl=%s, Act=%s, LoginRequired=%v, ACL=%d, isLogged=%v, isAjax=%v, route=%s, ReqUrl=%s", c.Ctx.Input.Method(), c.ctrlName, c.actName, c.isLoginRequired, c.acl, c.isLogged, c.IsAjax(), c.Data["RouterPattern"], c.Ctx.Request.URL.String())
		spew.Dump(c.Input()) // Input body


		//spew.Dump(c.Ctx.Request.Header.Get("User-Agent"))
	}

	// 접근 제한
	c.checkAccessPermission()

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
	c.Data["member"] = c.member
	c.Data["reqVars"] =c.Input()
}

func (c *baseController) loginRequired(required bool) {
	c.isLoginRequired = required
	if ! required {
		c.acl = 0
	}
}

func (c *baseController) checkLoginStatus() {
	val := c.GetSession("memberId")
	if val != nil {
		member, err := models.GetMember(map[string]interface{}{
			"t.member_id":val.(int),
		})
		checkErr(err)

		if member != nil {
			c.member = member
			c.member.Location, _ = time.LoadLocation(c.member.Timezone)
			c.isLogged = true
		}
	}
}

func (c *baseController) checkAccessPermission() {
	if c.isLoginRequired {
		if !c.isLogged {
			redirectUri := url.QueryEscape(c.Ctx.Request.RequestURI)
			c.Redirect("/signin?redirectUri="+redirectUri, 302)
		}

		if c.member.Position & c.acl < 1 {
			c.Abort("403")
		}
	}
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

func (c *baseController) audit(category string, message interface{}, detail interface{}) error {
	var memberId int
	if c.member != nil {
		memberId = c.member.MemberId
	}
	models.Audit(&objs.AuditMsg{memberId, "signin_failed", c.Ctx.Input.IP(), message, detail})
	return nil
}


func (c *baseController) toJson(logs interface{}, total int64, err error, fastPaging string) {
	if fastPaging == "on" {
		c.Data["json"] = logs
	} else {
		dbResult := objs.NewDbResult()
		if err == nil {
			dbResult.State = true
		}
		dbResult.Rows = logs
		dbResult.Total = total
		c.Data["json"] = dbResult
	}

	c.ServeJSON()
}