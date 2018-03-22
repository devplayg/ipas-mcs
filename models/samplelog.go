package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/libs"
	"github.com/devplayg/ipas-mcs/objs"
)

//var RegexFoundRows = regexp.MustCompile(`(?i)SELECT(\s+)SQL_CALC_FOUND_ROWS`)

func GetSamplelog(filter *objs.SampleFilter) ([]objs.SampleLog, int64, error) {
	var where string
	var rows []objs.SampleLog

	// 조건 설정
	args := make([]interface{}, 0)
	args = append(args, filter.StartDate+":00", filter.EndDate+":59")

	if len(filter.Org) > 0 {
		where += fmt.Sprintf(" and org in (%s)", libs.JoinInt(filter.Org, ","))
	}

	if len(filter.RiskLevel) > 0 {
		where += fmt.Sprintf(" and risk_level in (%s)", libs.JoinInt(filter.RiskLevel, ","))
	}

	if len(filter.Guid) > 0 {
		where += " and guid like ?"
		args = append(args, "%"+filter.Guid+"%")
	}

	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query
	query := `
		SELECT %s *
		from log_sample
		where date >= ? and date <= ? %s
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
