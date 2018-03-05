package models

//import (
//	"github.com/astaxie/beego/orm"
//)
//
//
//func GetSystemConfig() ([]SysConfig, error) {
//	query := "select section, keyword, value_s, value_n from sys_config"
//	rows := []SysConfig{}
//	o := orm.NewOrm()
//	_, err := o.Raw(query).QueryRows(&rows)
//
//	return rows, err
//}
//
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
