package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/ipas-mcs/libs"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
	"github.com/yl2chen/cidranger"
	"html/template"
	"net"
	"net/url"
	"strings"
	"time"
)

type CtrlPreparer interface {
	CtrlPrepare()
}
type LangPreparer interface {
	LangPrepare()
}

// 기본 Controller
type baseController struct {
	beego.Controller              // 메인 구조체 임베딩
	i18n.Locale                   // 다국어
	isLoginRequired  bool         // 로그인 필수 여부
	acl              uint         // 권한
	member           *objs.Member // 사용자 정보
	isLogged         bool         // 로그인 상태
	ctrlName         string       // Controller 이름
	actName          string       // Action 이름
	langMap map[string]string
	log *log.Logger
}

func (c *baseController) Prepare() {

	// Controller 와 Action 이름 설정
	c.ctrlName, c.actName = c.GetControllerAndAction()
	c.ctrlName = strings.ToLower(strings.TrimSuffix(c.ctrlName, "Controller"))

	// 기본 접근권한 설정
	c.isLoginRequired = true                   // 로그인 필수
	c.grant(objs.Superman, objs.Administrator) // 관리자 이상만 실행 허용
	c.isLogged = false                         // 로그인 상태

	// 호출된 Controller 접근권한 덮어쓰기
	if app, ok := c.AppController.(CtrlPreparer); ok {
		app.CtrlPrepare()
	}
	// 로그인 상태 체크
	c.checkLoginStatus()

	// 요청 디버깅 코드
	if beego.BConfig.RunMode == "dev" {

		log.Debug("=============START=================================")
		log.Debugf("Method=%s, Ctrl=%s, Act=%s, LoginRequired=%v, ACL=%d, isLogged=%v, isAjax=%v, Path=%s, Ext=%s, route=%s, ReqUrl=%s, remote_addr=%s",
			c.Ctx.Input.Method(),
			c.ctrlName,
			c.actName,
			c.isLoginRequired,
			c.acl,
			c.isLogged,
			c.IsAjax(),
			c.Ctx.Input.Param(":path"),
			c.Ctx.Input.Param(":ext"),
			c.Data["RouterPattern"],
			c.Ctx.Request.URL.String(),
			c.Ctx.Input.IP(),
		)
		spew.Dump(c.Input()) // Input body
		//spew.Dump(c.Ctx.Request.Header)
		//log.Debugf("Content-Type: %s", c.Ctx.Request.Header["Content-Type"])
		//spew.Dump(c.Ctx.Request.Header.Get("User-Agent"))
	}

	// 접근 제한
	c.checkAccessPermission()

	// 언어 설정
	c.setLangVer()
	c.loadFrontLang()
	//c.langMap = make(map[string]string)
	//c.langToFrontEnd()

	// 기본 템플릿 변수 설정
	c.Data["title"] = beego.BConfig.AppName
	//c.Data["member"] = c.GetMember()
	//c.Data["IsLogged"] = c.IsLogged
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["ctrl"] = c.ctrlName
	c.Data["act"] = c.actName
	c.Data["member"] = c.member
	c.Data["company_name"] = beego.AppConfig.DefaultString("company_name", "KYUNGWOO")
	c.Data["product_name"] = beego.AppConfig.DefaultString("product_name", "IPAS-MCS")
	c.Data["product_version"] = beego.AppConfig.DefaultString("product_version", "")
	c.Data["reqVars"] = c.Input()
}

func (c *baseController) loginRequired(required bool) {
	c.isLoginRequired = required
	if !required {
		c.acl = 0
	}
}

func (c *baseController) checkLoginStatus() {
	val := c.GetSession("memberId")
	if val != nil {
		member, err := models.GetMember(map[string]interface{}{
			"t.member_id": val,
		})
		if err != nil {
			log.Error(err)
		}

		if member != nil {
			c.member = member
			c.isLogged = true
			c.member.Location, _ = time.LoadLocation(c.member.Timezone)
		}
	}
}

func (c *baseController) checkAccessPermission() {
	if c.isLoginRequired { // 로그인이 요구돠는 페이지
		if !c.isLogged { // 로그인 되어 있지 않으면
			redirectUri := url.QueryEscape(c.Ctx.Request.RequestURI)
			c.Redirect("/signin?redirectUri="+redirectUri, 302) // 로그인화면으로 리다이렉션
		}

		if c.member.Position&c.acl < 1 { // 페이지 접근 권한이 없으면
			c.Abort("403")
		}

		// IP 접근 제어
		if c.Ctx.Input.IP() == "127.0.0.1" { // 로컬IP는 허용
			return
		}

		// 허가된 IP인지 확인
		allowedList := libs.SplitString(c.member.AllowedIp, ",")
		if len(allowedList) < 1 {
			c.Abort("403")
		}
		ranger := cidranger.NewPCTrieRanger()
		for _, s := range libs.SplitString(c.member.AllowedIp, ",") {
			_, network, _ := net.ParseCIDR(s)
			ranger.Insert(cidranger.NewBasicRangerEntry(*network))
		}
		containingNetworks, _ := ranger.ContainingNetworks(net.ParseIP(c.Ctx.Input.IP()))
		if len(containingNetworks) < 1 {
			c.Abort("403")
		}
	}
}

func (c *baseController) setTpl(tplName string) {
	c.TplName = c.ctrlName + "/" + tplName
}

func (c *baseController) grant(auth ...uint) {
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

// 감사로그
func (c *baseController) audit(category string, message interface{}, detail interface{}) error {
	var memberId int
	if c.member != nil {
		memberId = c.member.MemberId
	}
	err := models.Audit(&objs.AuditMsg{memberId, "signin_failed", c.Ctx.Input.IP(), message, detail})
	return err
}

func (c *baseController) loadFrontLang() {
	c.langMap = make(map[string]string)
	c.addToFrontLang("yes,no,msg.confirm_delete")
	if app, ok := c.AppController.(LangPreparer); ok {
		app.LangPrepare()
	}

	c.Data["frontLang"] = c.langMap
}

func (c *baseController) addToFrontLang(str string) {
	list := libs.SplitString(str, `[\s|,]+`)
	for _, r := range list {
		c.langMap[r] = c.Tr(r)
	}
}

func (c *baseController) serveResultJson(logs interface{}, total int64, err error, fastPaging string) {
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