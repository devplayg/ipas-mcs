package libs

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/devplayg/golibs/secureconfig"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// 초기화
func Initialize(processName string, encKey []byte, debug, verbose bool) {
	log.Debug("Initializing..")
	initLogger(processName, debug, verbose)

	if err := initDatabase(processName, encKey); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err := loadSystemConfig(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err := addExtraFunctions(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err := initLocale(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}


func loadSystemConfig() error {
	//rows, err := models.GetSystemConfig()
	//if err == nil {
	//	for _, r := range rows {
	//		m, ok := GlobalConfig[r.Section]
	//		if !ok {
	//			m = make(map[string]ConfigValue)
	//			GlobalConfig[r.Section] = m
	//		}
	//		m[r.Keyword] = ConfigValue{
	//			ValueS: r.ValueS,
	//			ValueN: r.ValueN,
	//		}
	//	}
	//}

	return nil
}

// 템플릿 변수 추가
func addExtraFunctions() error {
	// Extra functions
	if err := beego.AddFuncMap("literal", literal); err != nil {
		return err
	}

	if err := beego.AddFuncMap("i18n", i18n.Tr); err != nil {
		return err
	}

	return nil
}


// 다국어 설정
func initLocale() error {
	beego.Debug("Initializing locale..")
	langs := strings.Split("ko-kr|en-us|ja-jp", "|")
	names := strings.Split("KO|EN|JP", "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to load message file: " + err.Error())
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

//
func literal(text string) template.HTML {
	return template.HTML(text)
}
