package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
)

func GetMemberById(memberId int) (*objs.Member, error) {
	member := objs.Member{}

	o := orm.NewOrm()
	query := `
        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
        from mbr_member t
        	left outer join mbr_password t1 on t1.member_id = t.member_id
        where t.member_id = ?
    `
	err := o.Raw(query, memberId).QueryRow(&member)
	return &member, err
}