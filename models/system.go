package models

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"strings"
)

func GetSystemConfig() ([]objs.SysConfig, error) {
	query := "select section, keyword, value_s, value_n from sys_config"
	rows := []objs.SysConfig{}

	o := orm.NewOrm()
	_, err := o.Raw(query).QueryRows(&rows)

	return rows, err
}

func UpdateRow(tableName string, pkColumn string, pkValue interface{}, data map[string]interface{}) (sql.Result, error) {
	args := make([]interface{}, 0)
	setPhrases := make([]string, 0)
	for k, v := range data {
		setPhrases = append(setPhrases, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}
	args = append(args, pkValue)
	query := fmt.Sprintf("update %s set %s where %s = ?", tableName, strings.Join(setPhrases, ","), pkColumn)
	o := orm.NewOrm()
	return o.Raw(query, args).Exec()
}

// 감사로깅
func Audit(log *objs.AuditMsg) error {
	o := orm.NewOrm()

	var message string
	if log.Message != nil {
		m, _ := json.Marshal(log.Message)
		message = string(m)
	}

	// 간단한 감사이력 로깅
	query := "insert into adt_audit(member_id, category, ip, message) values(?, ?, inet_aton(?), ?)"
	rs, err := o.Raw(query, log.MemberId, log.Category, log.IP, message).Exec()
	if err != nil {
		o.Rollback()
		return err
	} else {
		defer o.Commit()

		// 상세 감사이력 로깅
		if log.Detail != nil {
			lastInsertId, _ := rs.LastInsertId()

			var detail string
			d, _ := json.Marshal(log.Detail)
			detail = string(d)
			query := "insert into adt_audit_detail(audit_id, detail) values(?, ?)"
			_, err2 := o.Raw(query, lastInsertId, detail).Exec()
			if err2 != nil {
				return err2
			}
		}
	}

	return nil
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
