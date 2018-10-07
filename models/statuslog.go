package models

import (
	"fmt"
	"github.com/devplayg/ipas-mcs/libs"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"time"
	"github.com/devplayg/ipas-server"
)

func GetIpasStatusLog(filter *objs.IpasFilter, member *objs.Member) ([]objs.IpasLog, int64, error) {
	var where string
	var rows []objs.IpasLog

	// 조건 설정
	args := make([]interface{}, 0)

	// 시간설정
	startDate, _ := time.ParseInLocation(ipasserver.DateDefault, filter.StartDate+":00", member.Location)
	endDate, _ := time.ParseInLocation(ipasserver.DateDefault, filter.EndDate+":59", member.Location)
	args = append(args, startDate.UTC().Format(ipasserver.DateDefault), endDate.UTC().Format(ipasserver.DateDefault))

	if member.Position < objs.Administrator {
		where += " and group_id in (select asset_id from mbr_asset where member_id = ?)"
		args = append(args, member.MemberId)
	}

	if len(filter.OrgId) > 0 {
		where += fmt.Sprintf(" and org_id in (%s)", libs.JoinInt(filter.OrgId, ","))
	}

	if len(filter.GroupId) > 0 {
		where += fmt.Sprintf(" and group_id in (%s)", libs.JoinInt(filter.GroupId, ","))
	}

	// 장비 태크 검색
	if len(filter.TagPattern) > 0 {
		where += " and (equip_id like ?)"
		cond := "%"+filter.TagPattern+"%"
		args = append(args, cond)
	}
	if len(filter.EquipId) > 0 {
		where += " and equip_id = ?"
		cond := filter.EquipId
		args = append(args, cond)
	}

	// 페이징 모드(고속/일반)
	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query
	query := `
		SELECT 	%s date, org_id, group_id, session_id, equip_id, latitude, longitude, speed
				, snr, usim, ip, recv_date
		from log_ipas_status
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

