package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"fmt"
	"strings"
)

func GetMember(condMap map[string]interface{}) (*objs.Member, error) {

	// 검색조건 생성
	args := make([]interface{}, 0)
	conditions := make([]string, 0)
	for k, v := range condMap {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}

	member := objs.Member{}
	query := `
        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
        from mbr_member t
        	left outer join mbr_password t1 on t1.member_id = t.member_id
        where %s
    `
	query = fmt.Sprintf(query, strings.Join(conditions, ","))
	o := orm.NewOrm()
	err := o.Raw(query, args).QueryRow(&member)
	return &member, err
}

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

func GetMemberByUsername(username string) (*objs.Member, error) {
	member := objs.Member{}
	o := orm.NewOrm()
	query := `
        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
        from mbr_member t
        	left outer join mbr_password t1 on t1.member_id = t.member_id
        where username = ?
    `
	err := o.Raw(query, username).QueryRow(&member)
	return &member, err
}


//func UpdateMember(memberId int, data map[string]interface{}) (sql.Result, error) {
//	query := `
//		update mbr_member
//		set %s
//		where member_id = ?
//	`
//	args := make([]interface{}, 0)
//	contents := make([]string, 0)
//	for k, v := range data {
//		contents = append(contents, fmt.Sprintf("%s = ?", k))
//		args = append(args, v)
//	}
//	args = append(args, memberId)
//	query = fmt.Sprintf(query, strings.Join(contents, ","))
//	o := orm.NewOrm()
//	return o.Raw(query, args).Exec()
//}
//
