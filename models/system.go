package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"database/sql"
	"fmt"
	"strings"
)


func GetSystemConfig() ([]objs.SysConfig, error) {
	query := "select section, keyword, value_s, value_n from sys_config"
	rows := []objs.SysConfig{}

	o := orm.NewOrm()
	_, err := o.Raw(query).QueryRows(&rows)

	return rows, err
}


func UpdateRow(tableName string, pkColumn string, pkValue interface{}, data map[string]interface{}) (sql.Result, error)  {
	args := make([]interface{}, 0)
	list := make([]string, 0)
	for k, v := range data {
		list = append(list, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}
	args = append(args, pkValue)
	query := fmt.Sprintf("update %s set %s where %s = ?", tableName, strings.Join(list, ","), pkColumn)
	o := orm.NewOrm()
	return o.Raw(query, args).Exec()
}

//func UpdateSystemConfig(configs []SysConfig) {
//	o := orm.NewOrm()
//	query := `
//        insert into sys_config(section, keyword, value_s, value_n)
//        values (?, ?, ?, ?)
//        on duplicate key update
//            value_s = values(value_s),
//            value_n = values(value_n);
//    `
//	stmt, _ := o.Raw(query).Prepare()
//	for _, v := range configs {
//		stmt.Exec(v.Section, v.Keyword, v.ValueS, v.ValueN)
//	}
//	stmt.Close()
//}
