package models

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

//
//import (
//	"fmt"
//	"github.com/astaxie/beego/orm"
//)

//select org_id, equip_id, group_id, equip_type, latitude, longitude, speed, snr, usim
//from ast_ipas
//
func GetIpaslist(filter *objs.IpasFilter) ([]objs.Ipas, int64, error) {

	args := make([]interface{}, 0)

	var where string
	var rows []objs.Ipas

	// 기관 또는 그룹 ID
	if len(filter.OrgId) > 0 {
		where += fmt.Sprintf(" and t.org_id in (%s)",JoinInt(filter.OrgId, ","))
	}
	if len(filter.GroupId) > 0 {
		where += fmt.Sprintf(" and t.group_id in (%s)",JoinInt(filter.GroupId, ","))
	}

	// 페이징 모드(고속/일반)
	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query

	query := `
		select %s org_id, equip_id, group_id, equip_type, latitude, longitude, speed, snr, usim, t.created, t.updated, t1.name org_name, t2.name group_name
		from ast_ipas t
			left outer join (
				select asset_id, name from ast_asset where class = 1 and type1 = 1
			)  t1 ON t1.asset_id = t.org_id
			left outer join (
				select asset_id, name from ast_asset where class = 1 and type1 = 2
			)  t2 ON t2.asset_id = t.group_id
		where true %s
		order by %s %s, equip_id asc
		limit ?, ?
	`
	query = fmt.Sprintf(query, filter.FoundRows, where, filter.Sort, filter.Order)
	args = append(args, filter.Offset, filter.Limit)

	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		spew.Dump(err)
	}

	if filter.FastPaging == "off" {
		if RegexFoundRows.MatchString(query) {
			dbResult := objs.NewDbResult()
			o.Raw("select FOUND_ROWS() total").QueryRow(dbResult)
			total = dbResult.Total
		}
	}
	return rows, total, err
}