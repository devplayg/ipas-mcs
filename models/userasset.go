package models

import (
	"github.com/devplayg/ipas-mcs/objs"
	"github.com/astaxie/beego/orm"
)
func GetUserassetsByClass(class int, member *objs.Member) ([]objs.Asset, error) {
	o := orm.NewOrm()

	var assets []objs.Asset
	var err error

	if member.Position >= objs.Administrator {
		query := `
			select  t.asset_id, class, parent_id, name, type1, type2, code,
					t.asset_id id,
					name text,
					type1 type
			from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
			where class & ? > 0
			order by class, name
		`
		_, err = o.Raw(query, class).QueryRows(&assets)
	} else {
		query := `
  			select  t.asset_id, class, parent_id, name, type1, type2, code,
					t.asset_id id,
					name text,
					type1 type
			from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
            where class & ? > 0 and (
				t.asset_id in (select asset_id from mbr_asset where member_id = ?)
				or
				t.asset_id in (
					select parent_id
					from ast_asset
					where asset_id in (select asset_id from mbr_asset where member_id = ?)
				)
			)
            order by class, name
		`
		_, err = o.Raw(query, class, member.MemberId, member.MemberId).QueryRows(&assets)
	}

	return assets, err
}
