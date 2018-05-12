package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
)

func GetStatsBy(member *objs.Member, filter *objs.StatsFilter) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	var assetId int
				args = append(args, "stats", "last_updated")

				if filter.OrgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
					if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
						assetId = -1
					} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
						assetId = member.MemberId * -1
					}
					args = append(args, assetId)
				} else { // 부분 자산통계 데이터 접근 시
					if filter.GroupId < 0 {
						if member.Position >= objs.Administrator {
							args = append(args, filter.GroupId)
						} else {
							// 일반 사용자는 기관 전체에 대한 통계데이터 접근 불가.
							// 개발 그룹에 대한 통계데이터는 접근 가능
							return nil, 0, nil
			}
		} else {
			if member.Position >= objs.Administrator {
				args = append(args, filter.GroupId)
			} else {
				where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
				args = append(args, filter.GroupId, member.MemberId)
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
	query = fmt.Sprintf(query, filter.StatsType, filter.AssetType, where)
	args = append(args, filter.Top)
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}


func GetStats(member *objs.Member,filter *objs.StatsFilter) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	var assetId int
	args = append(args, "stats", "last_updated")

	if filter.OrgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
			assetId = -1
		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
			assetId = member.MemberId * -1
		}
		args = append(args, assetId)
	} else { // 부분 자산통계 데이터 접근 시
		if filter.GroupId < 0 {
			if member.Position >= objs.Administrator {
				args = append(args, filter.OrgId)
			} else {
				// 일반 사용자는 기관 전체에 대한 통계데이터 접근 불가.
				// 개발 그룹에 대한 통계데이터는 접근 가능
				return nil, 0, nil
			}
		} else {
			if member.Position >= objs.Administrator {
				args = append(args, filter.GroupId)
			} else {
				where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
				args = append(args, filter.GroupId, member.MemberId)
			}
		}
	}

	query := `
		select date, asset_id, item, count, rank
		from stats_%s
		where date = (select value_s from sys_config where section = ? and keyword = ?) and asset_id = ? %s
		order by rank asc
		limit ?
	`
	// where date >= ? and date <= ? and asset_id = ? %s
	query = fmt.Sprintf(query, filter.StatsType, where)
	args = append(args, filter.Top)
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}
