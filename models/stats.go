package models

import (
	"github.com/devplayg/ipas-mcs/objs"
	"fmt"
	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
)

// 그룹통계(stats_*_by_group) 조회
func GetOrgGroupStats(member *objs.Member, p map[string]interface{}) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	args = append(args, p["from"], p["to"])

	if p["orgId"] == "0" && p["groupId"] == "-1" { // 전체 자산통계 요청 시
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
			args = append(args, -1)
		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
			args = append(args, member.MemberId * -1)
		}

	} else { // 특정 기관/그룹에 대한 통계 요청 시
		where += " and item = ?"
		args = append(args, p["orgId"], p["assetKey"])
		if member.Position < objs.Administrator { // 관리자 세션이 아니면, 허용된 자산통계데이터에 접근허용
			where += " and substring_index(item, '/', -1) in (select asset_id from mbr_asset where member_id = ?)"
			args = append(args, member.MemberId)
		}
	}

	args = append(args, p["top"])
	query := `
		select date, asset_id, item, count, rank
		from stats_%s_by_group
		where date >= ? and date <= ? and asset_id = ? %s 
		order by rank asc
		limit ?
	`
	query = fmt.Sprintf(query, p["statsType"], where)
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}

func GetStats(member *objs.Member, p map[string]interface{}) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	args = append(args, p["from"], p["to"])

	if p["assetId"].(int) <= 0 { // 전체 자산통계 요청 시
		if member.Position >= objs.Administrator { // 관리자 세션이면,
			p["assetId"] = -1
		} else {
			p["assetId"] = member.MemberId * -1
		}
		args = append(args, p["assetId"])

	} else { // 관리자 세션이 아니면, 허용된 자산통계데이터에 접근허용
		where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
		args = append(args, p["assetId"], member.MemberId)
	}

	query := `
		select date, asset_id, item, count, rank
		from stats_%s_by_%s
		where date >= ? and date <= ? and asset_id = ? %s 
		order by rank asc
		limit ?
	`
	args = append(args, p["top"])
	query = fmt.Sprintf(query, p["statsType"], p["assetType"], where)
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}
