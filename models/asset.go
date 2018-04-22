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
		query = `
            select  t.asset_id, class, parent_id, name, type1, type2, code,
                    t.asset_id id,
                    name text,
                    concat('type_', type1) type
            from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
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
	o := orm.NewOrm()
	var rs sql.Result
	var err error
	o.Begin()

	query := "insert into ast_asset(class, parent_id, name, type1, type2) values(?, ?, ?, ?, ?)"
	rs, err = o.Raw(query, asset.Class, asset.ParentId, asset.Name, asset.Type1, asset.Type2).Exec()
	if err != nil {
		o.Rollback()
		return nil, err
	}

	lastInsertId, err := rs.LastInsertId()
	if err != nil {
		o.Rollback()
		return nil, err
	}

	if asset.Class == 1 && asset.Type1 == 1 { // "기관" 등록이면
		query := `
			insert into ast_code(asset_id, code)
			values(?, ?)
			on duplicate key update code = values(code)
		`
		_, err = o.Raw(query, lastInsertId, asset.Code).Exec()
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}


	o.Commit()
	return rs, err

	//query := "insert into ast_asset(class, parent_id, name, type1, type2) values(?, ?, ?, ?, ?)"
	//
	//o := orm.NewOrm()
	//rs, err := o.Raw(query, asset.Class, asset.ParentId, asset.Name, asset.Type1, asset.Type2).Exec()
	//return rs, err
}

func GetAsset(assetId int) (objs.Asset, error) {
	var asset objs.Asset
	query := `
		select t.asset_id, class, parent_id, name, type1, type2, code
		from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
		where t.asset_id = ?
	`

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

func UpdateAsset(asset *objs.Asset) (sql.Result, error) {
	o := orm.NewOrm()
	var rs sql.Result
	var err error
	o.Begin()

	if len(asset.Code) > 0 {
		query := `
			insert into ast_code(asset_id, code)
			values(?, ?)
			on duplicate key update code = values(code)
		`
		_, err = o.Raw(query, asset.AssetId, asset.Code).Exec()
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}

	query := "update ast_asset set name = ? where asset_id = ?"
	rs, err = o.Raw(query, asset.Name, asset.AssetId).Exec()
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
