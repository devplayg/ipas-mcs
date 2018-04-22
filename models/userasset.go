package models
//
//import "github.com/astaxie/beego/orm"
//import (
//	"github.com/devplayg/ipas-mcs/objs"
//	"fmt"
//)
//func GetUserassetsByClass(class int, member *objs.Member) ([]objs.Asset, error) {
//	o := orm.NewOrm()
//
//	var assets []objs.Asset
//	var query string
//	var where string
//	var err error
//	args := make([]interface{}, 0)
//
//	if member.Position < objs.Administrator {
//		where += "select asset_id from mbr_asset where member_id = ?"
//		args = append(args, member.MemberId)
//	}
//
//	if class > 0 {
//		args := make([]interface{}, 0)
//		query := `
//            select  t.asset_id, class, parent_id, name, type1, type2, code,
//					t.asset_id id,
//					name text,
//					type1 type
//			from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
//            where class & ? > 0 %s
//            order by class, name
//        `
//        query = fmt.Sprintf(query, where)
//        //args = append([]interface{}, args)
//		_, err = o.Raw(query, args).QueryRows(&assets)
//	} else {
//		query = `
//            select  t.asset_id, class, parent_id, name, type1, type2, code,
//                    t.asset_id id,
//                    name text,
//                    concat('type_', type1) type
//            from ast_asset t left outer join ast_code t1 on t1.asset_id = t.asset_id
//			where true %s
//            order by class, name
//        `
//		query = fmt.Sprintf(query, where)
//		args = append(args, class, member.MemberId)
//		_, err = o.Raw(query, args).QueryRows(&assets)
//	}
//
//	return assets, err
//}
