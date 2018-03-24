package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
)

func GetAssetsByClass(class int) ([]objs.Asset, error) {
	o := orm.NewOrm()

	var assets []objs.Asset
	var query string
	var err error
	if class > 0 {
		query := `
            select  *,
                    concat('assetid_', asset_id) id,
                    name text,
                    type1 type
            from ast_asset
            where class & ? > 0
            order by class, name
        `
		_, err = o.Raw(query, class).QueryRows(&assets)
	} else {
		query = `
            select  *,
                    concat('assetid_', asset_id) id,
                    name text,
                    concat('type_', type1) type
            from ast_asset
            order by class, name
        `
		_, err = o.Raw(query).QueryRows(&assets)
	}

	return assets, err
}
//
////select distinct a.asset_id, a.class, a.parent_id, a.name, a.type1, a.type2, a.hostname, a.cidr, a.port
////from ast_asset a	join (
////	select t.asset_id, t.parent_id
////	from ast_asset t
////		join (
////			select asset_id, parent_id
////			from ast_asset
////			where asset_id in (select value_n from mbr_config where member_id = 11 and keyword = 'asset')
////    ) t1 on t1.asset_id = t.asset_id or t1.parent_id = t.asset_id
////) b on b.asset_id = a.asset_id or b.parent_id = a.asset_id
//
func GetAssetChildren(assetId int) ([]objs.Asset, error) {
	o := orm.NewOrm()

	var assets []objs.Asset
	query := `
        select  *
        from ast_asset
        where parent_id = ?
        order by class asc, name asc
    `
	_, err := o.Raw(query, assetId).QueryRows(&assets)
	return assets, err
}
//
//func UpdateAsset(asset Asset) (sql.Result, error) {
//	o := orm.NewOrm()
//	var where string
//	args := make([]interface{}, 0, 4)
//
//	query := `
//        update ast_asset
//        set asset_id = asset_id %s
//        where asset_id = ?;
//    `
//
//	// Name
//	if len(asset.Name) > 0 {
//		where += ", name = ?"
//		args = append(args, asset.Name)
//	}
//
//	// Asset ID
//	args = append(args, asset.AssetId)
//
//	// Update
//	query = fmt.Sprintf("update ast_asset set asset_id = asset_id %s where asset_id = ?", where)
//	return o.Raw(query, args).Exec()
//}
//
//func DeleteAsset(id_list []int) (err error) {
//	o := orm.NewOrm()
//
//	stmt1, _ := o.Raw("delete from ast_asset where asset_id = ?").Prepare()
//	stmt2, _ := o.Raw("delete from mbr_config where keyword = ? and value_n = ?").Prepare()
//
//	for _, assetId := range id_list {
//		_, e := stmt1.Exec(assetId)
//		if e != nil {
//			err = e
//		}
//
//		_, e = stmt2.Exec("asset", assetId)
//		if e != nil {
//			err = e
//		}
//	}
//	stmt1.Close()
//	stmt2.Close()
//
//	return err
//}
//
//func GetAssetIdListByMemberId(memberId, class int) ([]int, error) {
//	o := orm.NewOrm()
//
//	var list []int
//	args := make([]interface{}, 0, 4)
//	query := "select value_n from mbr_config where member_id = ? and keyword = ? and value_s = ?"
//
//	args = append(args, memberId, "asset", class)
//	_, err := o.Raw(query, args).QueryRows(&list)
//	return list, err
//}
//
