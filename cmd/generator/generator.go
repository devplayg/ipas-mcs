package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icrowley/fake"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Group struct {
	OrgId   int
	GroupId int
	Name    string
}

var (
	flags   *flag.FlagSet
	idPools orm.ParamsList
	groups  []Group
)

func init() {
	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&loc=%s%s",
		"root",
		"sniper123!@#",
		"tcp",
		"127.0.0.1",
		"3306",
		"ipasm",
		"utf8",
		strings.Replace("Asia/Seoul", "/", "%2F", -1),
		"&parseTime=true&allowAllFiles=true",
	)
	err := orm.RegisterDataBase("default", "mysql", connStr, 1, 1)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	query := "SELECT parent_id org_id, asset_id group_id, name FROM ast_asset where class =1 and type1 = 2"
	_, err = o.Raw(query).QueryRows(&groups)
	if err != nil {
		panic(err)
	}

	//spew.Dump(groups)
}

func main() {

	// 옵션
	flags = flag.NewFlagSet("", flag.ExitOnError)
	var (
		count = flags.Int("count", 10, "Count") // 버전
	)
	flags.Usage = printHelp // 도움말
	flags.Parse(os.Args[1:])

	generateIpasLogs(*count)
}

func generateIpasLogs(count int) {
	logs := make([]*objs.IpasLog, 0)
	for i := 0; i < count; i++ {
		logs = append(logs, newIpasLog())
	}

	// Create templ file
	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	for _, r := range logs {
		str := fmt.Sprintf("%s\t%d\t%d\t%s\t%s\t%d\t%d\t%3.6f\t%4.6f\t%d\t%d\t%d\t%d\t%d\n",
			r.Date.Format("2006-01-02 15:04:05"),
			r.OrgId,
			r.GroupId,
			r.EquipId,
			r.Target,
			r.SpeedingCount,
			r.ShockCount,
			r.Latitude,
			r.Longitude,
			r.WarningDist,
			r.CautionDist,
			r.V2vDist,
			r.ShockThreshold,
			r.SpeedThreshold,
		)
		tempFile.WriteString(str)
	}

	// Close
	tempFile.Close()

	// Insert
	query := `
		LOAD DATA LOCAL INFILE '%s' INTO TABLE log_ipas
		FIELDS TERMINATED BY '\t'
		LINES TERMINATED BY '\n'
		(date,org_id,group_id,equip_id,target,speeding_count,shock_count,latitude,longitude,warning_dist,caution_dist,v2v_dist,shock_threshold,speed_threshold)
	`
	query = fmt.Sprintf(query, filepath.ToSlash(tempFile.Name()))
	o := orm.NewOrm()
	rs, err := o.Raw(query).Exec()
	if err != nil {
		panic(err)
	}
	//
	affectedRows, _ := rs.RowsAffected()
	fmt.Printf("%d logs\n", affectedRows)
}

func randTag(name string) string {
	tagType := NumberRange(1, 3)
	prefix := ""

	if tagType == 1 {
		prefix = "VT_"
	} else if tagType == 2 {
		prefix = "ZT_"
	} else if tagType == 3 {
		prefix = "PT_"
	}
	prefix += name + "_"
	//prefix += idPools[NumberRange(0, len(idPools)-1)].(string)
	return prefix + fake.DigitsN(2)
}

func randTags(name string) []string {
	count := NumberRange(1, 4)
	list := make([]string, count)
	for i := 0; i < count; i++ {
		list[i] = randTag(name)
	}

	return list
}

func NumberRange(from, to int) int {
	return fake.Year(from-1, to)
}

func randGroup() Group {
	return groups[NumberRange(0, len(groups)-1)]
}

func newIpasLog() *objs.IpasLog {
	group := randGroup()
	return &objs.IpasLog{
		Date:           time.Now().Add(time.Duration(NumberRange(1, 60)) * time.Second),
		OrgId:          group.OrgId,
		GroupId:        group.GroupId,
		EquipId:        randTag(group.Name),
		Target:         strings.Join(randTags(group.Name), ","),
		SpeedingCount:  NumberRange(1, 10),
		ShockCount:     NumberRange(1, 10),
		Latitude:       fake.Latitude(),
		Longitude:      fake.Longitude(),
		WarningDist:    NumberRange(1, 10),
		CautionDist:    NumberRange(1, 10),
		V2vDist:        NumberRange(1, 10),
		ShockThreshold: NumberRange(1, 10),
		SpeedThreshold: NumberRange(1, 10),
	}
}

func printHelp() {
	fmt.Println(strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])))
	flags.PrintDefaults()
}
