package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
)

func UpdateSystemConfig(config []objs.SysConfig) error {
	o := orm.NewOrm()
	o.Begin()

	query := `
		insert into sys_config(section, keyword, value_s, value_n, updated)
		values(?, ?, ?, ?, now())
		on duplicate key update
			value_s = values(value_s),
			value_n = values(value_n),
			updated = values(updated)
	`

	p, err := o.Raw(query).Prepare()
	for _, c := range config {
		_, err := p.Exec(c.Section, c.Keyword, c.ValueS, c.ValueN)
		if err != nil {
			o.Rollback()
			return err
		}
	}
	p.Close()

	o.Commit()
	return err

}

//func GetSystemConfig() ([]objs.MultiValue, error) {
//	query := "select section, keyword, value_s, value_n from sys_config"
//	rows := []objs.MultiValue{}
//	o := orm.NewOrm()
//	_, err := o.Raw(query).QueryRows(&rows)
//
//	return rows, err
//}
