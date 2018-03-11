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
		conditions = append(conditions, fmt.Sprintf(" and %s = ?", k))
		args = append(args, v)
	}

	// 쿼리 생성
	query := `
        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
        from mbr_member t
        	left outer join mbr_password t1 on t1.member_id = t.member_id
        where true %s
    `
	query = fmt.Sprintf(query, strings.Join(conditions, ","))

	// 쿼리 실행
	o := orm.NewOrm()
	var member objs.Member
	err := o.Raw(query, args).QueryRow(&member)

	return &member, err
}

func GetMembers(filter *objs.CommonFilter) ([]objs.Member, int64, error) {
	var where string
	var rows []objs.Member

	// 조건 설정
	args := make([]interface{}, 0)

	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query
	query := `
        SELECT %s t.member_id, username, position, t.failed_login_count, t.status, timezone, t.name, t.session_id
        from mbr_member t
		where true %s
		order by %s %s
		limit ?, ?
	`
	query = fmt.Sprintf(query, filter.FoundRows, where, filter.Sort, filter.Order)
	args = append(args, filter.Offset, filter.Limit)

	o := orm.NewOrm()
	o.Begin()
	defer o.Commit()
	total, err := o.Raw(query, args).QueryRows(&rows)

	if filter.FastPaging == "off" {
		if RegexFoundRows.MatchString(query) {
			dbResult := objs.NewDbResult()
			o.Raw("select FOUND_ROWS() total").QueryRow(dbResult)
			total = dbResult.Total
		}
	}
	return rows, total, err
}


//func GetMembers(condMap map[string]interface{}) ([]objs.Member, error) {
//
//	// 검색조건 생성
//	var where string
//	args := make([]interface{}, 0)
//	for k, v := range condMap {
//		where += fmt.Sprintf(" and %s = ?", k)
//		args = append(args, v)
//	}
//
//	query := `
//        SELECT t.member_id, username, position, t.failed_login_count, t.status, timezone, t.name, t.session_id
//        from mbr_member t
//        where true %s
//	`
//	query = fmt.Sprintf(query, where)
//	o := orm.NewOrm()
//	var members []objs.Member
//	_, err := o.Raw(query, args).QueryRows(&members)
//	return members, err
//}



//func GetMemberById(memberId int) (*objs.Member, error) {
//	member := objs.Member{}
//
//	o := orm.NewOrm()
//	query := `
//        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
//        from mbr_member t
//        	left outer join mbr_password t1 on t1.member_id = t.member_id
//        where t.member_id = ?
//    `
//	err := o.Raw(query, memberId).QueryRow(&member)
//	return &member, err
//}
//
//func GetMemberByUsername(username string) (*objs.Member, error) {
//	member := objs.Member{}
//	o := orm.NewOrm()
//	query := `
//        SELECT t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status, timezone, t.name, t.session_id
//        from mbr_member t
//        	left outer join mbr_password t1 on t1.member_id = t.member_id
//        where username = ?
//    `
//	err := o.Raw(query, username).QueryRow(&member)
//	return &member, err
//}


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
