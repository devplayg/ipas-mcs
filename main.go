package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/devplayg/golibs/secureconfig"
	_ "github.com/devplayg/ipas-mcs/routers"
	"os"
	"path/filepath"
	"strings"
	"github.com/devplayg/ipas-mcs/libs"
)

var (
	flags *flag.FlagSet
)

func main() {

	// 옵션
	flags = flag.NewFlagSet("", flag.ExitOnError)
	var (
		version   = flags.Bool("version", false, "Version")            // 버전
		setConfig = flags.Bool("config", false, "Edit configurations") // 환경설정
		debug     = flags.Bool("debug", false, "Debug")                // 디버그
		verbose   = flags.Bool("v", false, "Verbose")
	)
	flags.Usage = printHelp // 도움말
	flags.Parse(os.Args[1:])

	// 버전 표시
	if *version {
		version := beego.AppConfig.DefaultString("app::version", "unknown")
		fmt.Printf("%s v%s\n", beego.BConfig.AppName, version)
		return
	}

	// 환경설정
	processName := getProcessName()
	encKey := getEncryptionKey()
	if *setConfig {
		keys := "db.hostname, db.port, db.username, db.password, db.database"
		secureconfig.SetConfig(processName+".enc", keys, encKey[:])
		return
	}

	// 초기화
	libs.Initialize(processName, encKey, *debug, *verbose)

	// 시작
	beego.Run()
}

func printHelp() {
	fmt.Println(strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])))
	flags.PrintDefaults()
}

func getProcessName() string {
	return strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
}

func getEncryptionKey() []byte {
	key := sha256.Sum256([]byte("Charlie Hunter - No Woman No Cry"))
	return key[:]
}


