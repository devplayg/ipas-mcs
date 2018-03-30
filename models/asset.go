package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"strings"
)

func GetAssetsByClass(class int) ([]objs.Asset, error) {
	o := orm.NewOrm()

	var assets []objs.Asset
	var query string
	var err error
	if class > 0 {
		query := `
            select  asset_id, class, parent_id, name, type1, type2,
                    asset_id id,
                    name text,
                    type1 type
            from ast_asset
            where class & ? > 0
            order by class, name
        `
		_, err = o.Raw(query, class).QueryRows(&assets)
	} else {
		query = `
            select  asset_id, class, parent_id, name, type1, type2,
                    asset_id id,
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
//func GetAssetChildren(assetId int) ([]objs.Asset, error) {
//	o := orm.NewOrm()
//
//	var assets []objs.Asset
//	query := `
//        select  *
//        from ast_asset
//        where parent_id = ?
//        order by class asc, name asc
//    `
//	_, err := o.Raw(query, assetId).QueryRows(&assets)
//	return assets, err
//}

func AddAsset(asset objs.Asset) (sql.Result, error) {
	query := "insert into ast_asset(class, parent_id, name, type1, type2) values(?, ?, ?, ?, ?)"

	o := orm.NewOrm()
	rs, err := o.Raw(query, asset.Class, asset.ParentId, asset.Name, asset.Type1, asset.Type2).Exec()
	return rs, err
}

func GetAsset(assetId int) (objs.Asset, error) {
	var asset objs.Asset
	query := "select asset_id, class, parent_id, name, type1, type2 from ast_asset where asset_id = ?"

	o := orm.NewOrm()
	err := o.Raw(query, assetId).QueryRow(&asset)
	return asset, err
}

func RemoveAsset(assetIdList []int) (sql.Result, error) {
	o := orm.NewOrm()
	var rs sql.Result
	var err error
	o.Begin()

	query := "delete from ast_asset where %s in (%s)"
	query = fmt.Sprintf(query, "parent_id", JoinInt(assetIdList, ","))
	_, err = o.Raw(query).Exec()
	if err != nil {
		o.Rollback()
		return nil, err
	}

	query = "delete from ast_asset where %s in (%s)"
	query = fmt.Sprintf(query, "asset_id", JoinInt(assetIdList, ","))
	rs, err = o.Raw(query).Exec()
	if err != nil {
		o.Rollback()
		return nil, err
	}

	o.Commit()
	return rs, err
}

func JoinInt(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
