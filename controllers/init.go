package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/devplayg/golibs/secureconfig"
	"github.com/devplayg/ipas-mcs/models"
	"github.com/devplayg/ipas-mcs/objs"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Multi-language
type langType struct {
	Lang, Name string
}

var langTypes []*langType // Languages are supported.
var assetMap sync.Map

// 초기화
func Initialize(processName string, encKey []byte, debug, verbose bool) {
	initLogger(processName, debug, verbose)
	log.Info("Initializing..")

	// 프레임워크 초기화
	initFramework()

	// 데이터베이스 초기화
	if err := initDatabase(processName, encKey); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 시스템 환경변수 초기화
	if err := loadGlobalConfig(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 템블릿 함수 추가
	if err := addExtraFunctions(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 다국어 기능 초기화
	if err := initLocale(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// 다국어 기능 초기화
	if err := loadAssets(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func loadAssets() error {
	assets, err := models.GetAssetsByClass(RootId)
	if err != nil {
		return err
	}
	for _, a := range assets {
		assetMap.Store(a.AssetId, a)
	}

	return nil
}

func initFramework() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 60 * 60 * 24 * 7
	beego.BConfig.WebConfig.Session.SessionName = "ipmSessionID"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
}

func loadGlobalConfig() error {

	// 기본값  설정
	objs.GlobalConfig.Store("login_allow_multiple_login", objs.MultiValue{"on", 0})
	objs.GlobalConfig.Store("login_block_seconds", objs.MultiValue{"", 60})
	objs.GlobalConfig.Store("login_max_failed_login_attempts", objs.MultiValue{"", 5})
	objs.GlobalConfig.Store("system_data_retention_days", objs.MultiValue{"", 365})
	objs.GlobalConfig.Store("system_use_namecard", objs.MultiValue{"on", 0})

	// DB 값 조회
	rows, err := models.GetSystemConfig()
	if err != nil {
		return err
	}

	for _, r := range rows {
		objs.GlobalConfig.Store(
			r.Section+"_"+r.Keyword,
			objs.MultiValue{r.ValueS, r.ValueN},
		)
	}

	return nil
}

// 템플릿 변수 추가
func addExtraFunctions() error {
	// 순수 문자열 출력 함수
	if err := beego.AddFuncMap("literal", literal); err != nil {
		return err
	}

	if err := beego.AddFuncMap("toLower", toLower); err != nil {
		return err
	}

	// 다국어 지원 함수
	if err := beego.AddFuncMap("i18n", i18n.Tr); err != nil {
		return err
	}

	return nil
}

// 다국어 설정
func initLocale() error {
	log.Debug("Initializing locale..")
	//languages := strings.Split("ko-kr|en-us|ja-jp", "|")
	//names := strings.Split("KO|EN|JP", "|")
	languages := strings.Split("ko-kr|en-us", "|")
	names := strings.Split("KO|EN", "|")
	langTypes = make([]*langType, 0, len(languages))
	for i, v := range languages {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range languages {
		log.Debug("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			log.Error("Fail to load message file: " + err.Error())
		}
	}

	return nil
}

// 데이터베이스 초기화
func initDatabase(processName string, encKey []byte) error {
	conf, err := secureconfig.GetConfig(processName+".enc", encKey)
	if err != nil {
		return err
	}

	maxIdle := beego.AppConfig.DefaultInt("db_master::maxidle", 3)
	maxConn := beego.AppConfig.DefaultInt("db_master::maxopen", 3)
	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&loc=%s%s",
		conf["db.username"],
		conf["db.password"],
		"tcp",
		conf["db.hostname"],
		conf["db.port"],
		conf["db.database"],
		"utf8",
		strings.Replace(beego.AppConfig.DefaultString("timezone", "Asia/Seoul"), "/", "%2F", -1),
		"&parseTime=true",
	)
	log.Debug("Connection string:", connStr)
	log.Debug("Max idle connections:", maxIdle)
	log.Debug("Max open connections:", maxConn)
	err = orm.RegisterDataBase("default", "mysql", connStr, maxIdle, maxConn)
	if err != nil {
		return err
	}

	return nil
}

// 로깅 초기화
func initLogger(processName string, debug, verbose bool) {
	// Set log format
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: true,
	})

	// Set log level
	if debug {
		log.SetLevel(log.DebugLevel)
		orm.Debug = false
	}

	if verbose {
		log.SetOutput(os.Stdout)
		orm.DebugLog = orm.NewLog(os.Stdout)
	} else {
		var logFile string
		if debug {
			logFile = filepath.Join(filepath.Dir(os.Args[0]), processName+"-debug.log")
			os.Remove(logFile)

		} else {
			logFile = filepath.Join(filepath.Dir(os.Args[0]), processName+".log")
		}

		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			log.SetOutput(file)
			orm.DebugLog = orm.NewLog(file)
		} else {
			log.SetOutput(os.Stdout)
			orm.DebugLog = orm.NewLog(os.Stdout)
			log.Error(err)
		}
	}

	if log.GetLevel() != log.InfoLevel {
		log.Infof("LoggingLevel=%s", log.GetLevel())
	}

}

//
func literal(s string) template.HTML {
	return template.HTML(s)
}

func toLower(s string) template.HTML {
	return template.HTML(strings.ToLower(s))
}

func CheckError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func init() {
	//rand.Seed(time.Now().UnixNano())
}
