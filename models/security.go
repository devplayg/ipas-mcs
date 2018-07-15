package models

import (
	"github.com/devplayg/ipas-mcs/objs"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetSecurityLogs(filter objs.SecurityLogFilter) ([]objs.SecurityLog, int64, error) {
	var where string
	var rows []objs.SecurityLog

	// 조건 설정
	args := make([]interface{}, 0)
	args = append(args, filter.StartDate+":00", filter.EndDate+":59")

	// 페이징 모드(고속/일반)
	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query
	query := `
		select %s t.audit_id, t.date, t.member_id, category, ip, message, t1.detail, m.username
		from adt_audit t 
			left outer join adt_audit_detail t1 on t1.audit_id = t.audit_id
			left outer join mbr_member m on m.member_id = t.member_id
		where t.date >= ? and t.date <= ? %s
		order by %s %s
		limit ?, ?
	`
	query = fmt.Sprintf(query, filter.FoundRows, where, filter.Sort, filter.Order)
	args = append(args, filter.Offset, filter.Limit)
	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	total, err := o.Raw(query, args).QueryRows(&rows)

	if filter.FastPaging == "off" {
		if RegexFoundRows.MatchString(query) {
			dbResult := objs.NewDbResult()
			o.Raw("select FOUND_ROWS() total").QueryRow(dbResult)
			total = dbResult.Total
		}
	}
	return rows, total, err
}

