package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/devplayg/ngclient/models"
	"github.com/devplayg/golibs/secureconfig"
	"strings"
	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"fmt"
)

type CtrlPreparer interface {
	CtrlPrepare()
}

type baseController struct {
	beego.Controller               // 메인 구조체 임베딩
	i18n.Locale                    // 다국어
	isLoginRequired bool          // 로그인 필수 여부
	acl              int           // 권한
	member           models.Member // 사용자 정보
	isLogged         bool          // 로그인 상태
	ctrlName         string        // Controller 이름
	actName          string        // Action 이름
}

func (c *baseController) Prepare() {
	// 기본권한 설정
	c.isLoginRequired = true        // 로그인 필수
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

// 초기화
func Initialize(processName string, encKey []byte, debug, verbose bool) {
	log.Debug("Initializing..")
	initLogger(processName, debug, verbose)

	if err := initDatabase(processName, encKey); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

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
		strings.Replace(beego.AppConfig.DefaultString("timezone", "Asia/Seoul"), "/", "%2F", -1), "&parseTime=true",
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
		orm.Debug = true
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
		}
	}

	if log.GetLevel() != log.InfoLevel {
		log.Infof("LoggingLevel=%s", log.GetLevel())
	}
}