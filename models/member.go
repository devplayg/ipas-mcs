package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
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

func GetMembers(filter *objs.PagingFilter) ([]objs.Member, int64, error) {
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

func AddMember(m *objs.Member) (sql.Result, error) {
	o := orm.NewOrm()
	o.Begin()

	// 사용자 정보 등록
	query := `
		insert into mbr_member(username, name, email, position)
		values(?, ?, ?, ?);
	`
	rs, err := o.Raw(query, m.Username, m.Name, m.Email, m.Position).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	// 비밀번호 등록
	lastInsertId, _ := rs.LastInsertId()
	query = "insert into mbr_password(member_id, password) values(?, ?)"
	rs, err = o.Raw(query, lastInsertId, m.PasswordConfirm).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	o.Commit()
	return rs, err
}

// Add , Update / Remove /
//func
