package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"strings"
)

func GetAllSystemConfig() ([]objs.SysConfig, error) {
	query := "select section, keyword, value_s, value_n from sys_config"
	var rows []objs.SysConfig

	o := orm.NewOrm()
	_, err := o.Raw(query).QueryRows(&rows)

	return rows, err
}

func GetSystemConfig(section, keyword string) ([]objs.SysConfig, error) {
	query := "select section, keyword, value_s, value_n from sys_config where section = ?"
	var where string
	args := make([]interface{}, 0)
	args = append(args, section)

	if len(keyword) > 0 {
		where += " and keyword = ?"
		args = append(args, keyword)
	}
	var rows []objs.SysConfig
	o := orm.NewOrm()
	_, err := o.Raw(query + where, args).QueryRows(&rows)

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

func RemoveRow(tableName string, pkColumn string, pkValue interface{}) (sql.Result, error) {
	query := fmt.Sprintf("delete from %s where %s = ?", tableName, pkColumn)
	o := orm.NewOrm()
	return o.Raw(query, pkValue).Exec()
}

// 감사로깅
func Audit(log *objs.AuditMsg) error {
	o := orm.NewOrm()
	o.Begin()

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

func GetServer(s objs.Server) (*objs.Server, error ){
	query := "select * from ast_server where true"
	var where string
	args := make([]interface{}, 0)

	if s.ID > 0 {
		where += " and server_id = ?"
		args = append(args, s.ID)
	}

	if s.Category1 > 0 {
		where += " and category1 = ?"
		args = append(args, s.Category1)
	}

	if s.Category2 > 0 {
		where += " and category2 = ?"
		args = append(args, s.Category2)
	}

	var server objs.Server

	o := orm.NewOrm()
	err := o.Raw(query + where, args).QueryRow(&server)
	return &server, err
}