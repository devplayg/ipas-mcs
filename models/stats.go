package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
)

func GetStats(member *objs.Member, orgId, groupId int, p map[string]interface{}) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	args = append(args, "stats", "last_updated")

	if orgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
			p["assetId"] = -1
		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
			p["assetId"] = member.MemberId * -1
		}
		args = append(args, p["assetId"])
	} else { // 부분 자산통계 데이터 접근 시
		if groupId < 0 {
			if member.Position >= objs.Administrator {
				args = append(args, orgId)
			} else {
				// 일반 사용자는 기관 전체에 대한 통계데이터 접근 불가.
				// 개발 그룹에 대한 통계데이터는 접근 가능
				return nil, 0, nil
			}
		} else {
			if member.Position >= objs.Administrator {
				args = append(args, groupId)
			} else {
				where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
				args = append(args, groupId, member.MemberId)
			}
		}
	}

	query := `
		select date, asset_id, item, count, rank
		from stats_%s_by_%s
		where date = (select value_s from sys_config where section = ? and keyword = ?) and asset_id = ? %s
		order by rank asc
		limit ?
	`
	// where date >= ? and date <= ? and asset_id = ? %s
	query = fmt.Sprintf(query, p["statsType"], p["assetType"], where)
	args = append(args, p["top"])
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}

//func GetGroupStats(member *objs.Member, orgId, groupId int, p map[string]interface{}) ([]objs.Stats, int64, error) {
	//var where string
	//var args []interface{}
	//args = append(args, p["from"], p["to"])
	//
	//if orgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
	//	if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
	//		p["assetId"] = -1
	//	} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
	//		p["assetId"] = member.MemberId * -1
	//	}
	//	args = append(args, p["assetId"])
	//} else { // 부분 자산통계 데이터 접근 시
	//	if groupId < 0 {
	//		if member.Position >= objs.Administrator {
	//			args = append(args, orgId)
	//		} else {
	//			// 일반 사용자는 기관 전체에 대한 통계데이터 접근 불가.
	//			// 개발 그룹에 대한 통계데이터는 접근 가능
	//			return nil, 0, nil
	//		}
	//	} else {
	//		if member.Position >= objs.Administrator {
	//			args = append(args, groupId)
	//		} else {
	//			where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
	//			args = append(args, groupId, member.MemberId)
	//		}
	//	}
	//}
	//
	//query := `
	//	select date, asset_id, item, count, rank
	//	from stats_%s_by_%s
	//	where date >= ? and date <= ? and asset_id = ? %s
	//	order by rank asc
	//	limit ?
	//`
	//query = fmt.Sprintf(query, p["statsType"], p["assetType"], where)
	//args = append(args, p["top"])
	//var rows []objs.Stats
	//o := orm.NewOrm()
	//total, err := o.Raw(query, args).QueryRows(&rows)
	//if err != nil {
	//	log.Error(err)
	//}
	//return rows, total, err

	//var where string
	//var args []interface{}
	//args = append(args, p["from"], p["to"])

	//	if p["orgId"] == "0" && p["groupId"] == "-1" { // 전체 자산통계 요청 시
	//		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
	//			args = append(args, -1)
	//		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
	//			args = append(args, member.MemberId * -1)
	//		}
	//
	//	} else { // 특정 기관/그룹에 대한 통계 요청 시
	//		where += " and item = ?"
	//		args = append(args, p["orgId"], p["assetKey"])
	//		if member.Position < objs.Administrator { // 관리자 세션이 아니면, 허용된 자산통계데이터에 접근허용
	//			where += " and substring_index(item, '/', -1) in (select asset_id from mbr_asset where member_id = ?)"
	//			args = append(args, member.MemberId)
	//		}
	//	}


	//args = append(args, p["top"])
	//query := `
	//	select date, asset_id, item, count, rank
	//	from stats_%s_by_group
	//	where date >= ? and date <= ? and asset_id = ? %s
	//	order by rank asc
	//	limit ?
	//`
	//query = fmt.Sprintf(query, p["statsType"], where)
	//var rows []objs.Stats
	//o := orm.NewOrm()
	//total, err := o.Raw(query, args).QueryRows(&rows)
	//if err != nil {
	//	log.Error(err)
	//}
	//return rows, total, err
//}

//
//// 그룹통계(stats_*_by_group) 조회
//func GetOrgGroupStats(member *objs.Member, p map[string]interface{}) ([]objs.Stats, int64, error) {
//	var where string
//	var args []interface{}
//	args = append(args, p["from"], p["to"])
//
//	if p["orgId"] == "0" && p["groupId"] == "-1" { // 전체 자산통계 요청 시
//		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
//			args = append(args, -1)
//		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
//			args = append(args, member.MemberId * -1)
//		}
//
//	} else { // 특정 기관/그룹에 대한 통계 요청 시
//		where += " and item = ?"
//		args = append(args, p["orgId"], p["assetKey"])
//		if member.Position < objs.Administrator { // 관리자 세션이 아니면, 허용된 자산통계데이터에 접근허용
//			where += " and substring_index(item, '/', -1) in (select asset_id from mbr_asset where member_id = ?)"
//			args = append(args, member.MemberId)
//		}
//	}
//
//	args = append(args, p["top"])
//	query := `
//		select date, asset_id, item, count, rank
//		from stats_%s_by_group
//		where date >= ? and date <= ? and asset_id = ? %s
//		order by rank asc
//		limit ?
//	`
//	query = fmt.Sprintf(query, p["statsType"], where)
//	var rows []objs.Stats
//	o := orm.NewOrm()
//	total, err := o.Raw(query, args).QueryRows(&rows)
//	if err != nil {
//		log.Error(err)
//	}
//	return rows, total, err
//}
//
//func GetStats(member *objs.Member, p map[string]interface{}) ([]objs.Stats, int64, error) {
//	var where string
//	var args []interface{}
//	args = append(args, p["from"], p["to"])
//
//	assetId, _ := strconv.Atoi(p["orgId"].(string))
//	if assetId <= 0 { // 전체 자산통계 요청 시
//		if member.Position >= objs.Administrator { // 관리자 세션이면,
//			p["assetId"] = -1
//		} else {
//			p["assetId"] = member.MemberId * -1
//		}
//		args = append(args, p["assetId"])
//
//	} else { // 관리자 세션이 아니면, 허용된 자산통계데이터에 접근허용
//		where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
//		args = append(args, assetId, member.MemberId)
//	}
//
//	query := `
//		select date, asset_id, item, count, rank
//		from stats_%s_by_%s
//		where date >= ? and date <= ? and asset_id = ? %s
//		order by rank asc
//		limit ?
//	`
//	args = append(args, p["top"])
//	query = fmt.Sprintf(query, p["statsType"], p["assetType"], where)
//	var rows []objs.Stats
//	o := orm.NewOrm()
//	total, err := o.Raw(query, args).QueryRows(&rows)
//	if err != nil {
//		log.Error(err)
//	}
//	return rows, total, err
//}
