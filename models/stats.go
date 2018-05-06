package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"fmt"
)


func GetStats(p map[string]interface{})  ([]objs.Stats, int64, error) {
	query := `
		select date, asset_id, item, count, rank
		from stats_%s_by_%s
		where date >= ? and date <= ? and asset_id = ? 
		order by rank asc
		limit ?
	`
	query = fmt.Sprintf(query, p["statsType"], p["assetType"])
	var rows []objs.Stats
	o := orm.NewOrm()
	total, err := o.Raw(query,  p["from"], p["to"], p["assetId"], p["top"]).QueryRows(&rows)
	return rows,	 total, err
}
