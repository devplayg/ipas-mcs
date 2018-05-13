package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/ipas-mcs/objs"
	log "github.com/sirupsen/logrus"
)

// 통계 검색 by 기관, 그룹, 장비(stats_evt?_by_(group|equip)
func GetStatsBy(member *objs.Member, filter *objs.StatsFilter) ([]objs.Stats, int64, error) {
	var where string
	var args []interface{}
	var assetId int
	query := `
        select date, asset_id, item, count, rank
        from stats_%s_by_%s
        where date = (select value_s from sys_config where section = ? and keyword = ?) and asset_id = ? %s
        order by rank asc
        limit ?
    `
	args = append(args, "stats", "last_updated")

	if filter.OrgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
			assetId = -1
		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
			assetId = member.MemberId * -1
		}
		args = append(args, assetId)
	} else { // 부분 자산통계 데이터 접근 시
		if filter.GroupId < 0 { // 전체 그룹 데이터 접근 시
			if member.Position >= objs.Administrator {
				args = append(args, filter.GroupId)
			} else {
				args = append(args, filter.OrgId)
				if filter.AssetType == "group" {
					where += " and SUBSTRING_INDEX(item, '/', -1) in (select asset_id from mbr_asset where member_id = ?)"
					args = append(args, member.MemberId)
				} else if filter.AssetType == "equip" {
					where += " and item in (select equip_id from ast_ipas where org_id = ? and group_id in (select asset_id from mbr_asset where member_id = ?))"
					args = append(args, filter.OrgId, member.MemberId)
				} else {
					return nil, 0, nil
				}
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

// stats_evt
func GetStats(member *objs.Member, filter *objs.StatsFilter) ([]objs.Stats, int64, error) {
	spew.Dump(filter)
	var where string
	var args []interface{}
	var assetId int
	query := `
        select date, asset_id, item, count, rank
        from stats_%s
        where date = (select value_s from sys_config where section = ? and keyword = ?) %s
        order by rank asc
        limit ?
    `
	args = append(args, "stats", "last_updated")

	if filter.OrgId < 1 { // 전체 자산통계 데이터 요청 시(org id가 -1인 경우. 참고: 0인 경우는 없음)
		where += " and asset_id = ?"
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 대한 접근 가능
			assetId = -1
		} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
			assetId = member.MemberId * -1
		}
		args = append(args, assetId)
	} else { // 부분 자산통계 데이터 접근 시
		if filter.GroupId < 0 { // 모든 그룹에 대한 조회 요청 시(기관만 선택되고, 그룹은 선택되지 않은 경우)
			if member.Position >= objs.Administrator { // 관리자 세션이면
				where += " and asset_id = ?"
				args = append(args, filter.OrgId)
			} else { // 일반사용자 세션이면, 허용된 통계데이터에만 접근 가능
				// 쿼리 다시 쓰기
				query = `
					select date, asset_id, item, count, rank
                    from stats_%s
                    where date = (select value_s from sys_config where section = ? and keyword = ?) 
                        and asset_id in (
                            select asset_id 
                            from ast_asset 
                            where parent_id = ? and asset_id in (select asset_id from mbr_asset where member_id = ?)
                        ) %s
                    order by count desc, item asc
                    limit ?
                `
				args = append(args, filter.OrgId, member.MemberId)
			}
		} else {
			where += " and asset_id = ?"
			if member.Position >= objs.Administrator { // 관리자 세션이면
				args = append(args, filter.GroupId)
			} else {
				where += " and asset_id in (select asset_id from mbr_asset where member_id = ?)"
				args = append(args, filter.GroupId, member.MemberId)
			}
		}
	}

	// where date >= ? and date <= ? and asset_id = ? %s
	query = fmt.Sprintf(query, filter.StatsType, where)
	args = append(args, filter.Top)

	if filter.StatsType == "evt" {
		spew.Dump(query)
		spew.Dump(args)
	}
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query, args).QueryRows(&rows)
	if err != nil {
		log.Error(err)
	}
	return rows, total, err
}

func GetEquipCountByType(member *objs.Member, filter *objs.StatsFilter) ([]objs.TagCount, error) {
	var where string
	var args []interface{}
	query := `
        select equip_type, count(*) count
        from ast_ipas
        where true %s
        group by equip_type
    `
	if filter.OrgId < 1 { // 전체통계 요청 시 (OrgId 가 -1인 경우에 해당. OrgId 가 0인 경우는 없음)
		if member.Position >= objs.Administrator { // 관리자 세션이면 전체통계(asset_id=-1)에 접근 가능
		} else { // 일반사용자 세션이면 접근제어
		}
	} else { // 기관통계 요청 시
		if filter.GroupId < 0 { // 기관 전체통계 요청 시 
			if member.Position >= objs.Administrator { // 관리자 세션이면 허용
            } else { // 일반사용자 세션이면 접근제어
            }
		} else { // 특정 그룹통계 요청 시
			if member.Position >= objs.Administrator { // 관리자 세션이면 허용
			} else { // 일반사용자 세션이면 접근제어
			}
		}
	}

	var rows []objs.TagCount
	o := orm.NewOrm()
	query = fmt.Sprintf(query, where)
	_, err := o.Raw(query, args).QueryRows(&rows)
	return rows, err

}
