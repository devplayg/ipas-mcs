package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/objs"
	"strings"
)

// 단일 사용자 검색
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
		select 	t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status,
				timezone, t.name, t.session_id, group_concat(inet_ntoa(t2.ip), '/', t2.cidr) allowed_ip, email,
				last_success_login, last_failed_login
		from mbr_member t
			left outer join mbr_password t1 on t1.member_id = t.member_id
			left outer join mbr_allowed_ip t2 on t2.member_id =  t.member_id
		where true %s
		group by member_id
	`
	query = fmt.Sprintf(query, strings.Join(conditions, ","))

	// 쿼리 실행
	o := orm.NewOrm()
	var member objs.Member
	err := o.Raw(query, args).QueryRow(&member)

	return &member, err
}

// 다중 사용자 검색
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
		select %s t.member_id, t.username, t.position, t1.password, t1.salt, t.failed_login_count, t.status,
				timezone, t.name, t.session_id, group_concat(inet_ntoa(t2.ip), '/', t2.cidr) allowed_ip, email,
				last_success_login, last_failed_login, t3.asset_summary
		from mbr_member t
			left outer join mbr_password t1 on t1.member_id = t.member_id
			left outer join mbr_allowed_ip t2 on t2.member_id =  t.member_id
			left outer join (
				select x.member_id, group_concat(x1.name, '__--__', x.name separator '__//__') asset_summary
				from (
					select o.member_id, o.asset_id, o1.parent_id, o1.name
					from mbr_asset o left outer join ast_asset o1 on o1.asset_id = o.asset_id
					order by parent_id, o1.name
				) x left outer join ast_asset x1 on x1.asset_id =  x.parent_id
				group by member_id
			) t3 on t3.member_id = t.member_id
		where true %s
		group by member_id
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

	// 접속 허용 IP 등록(prepared statement)
	p, err := o.Raw("insert into mbr_allowed_ip(member_id, ip, cidr) values(?, inet_aton(?), ?)").Prepare()
	for _, v := range m.AllowedIpList {
		_, err := p.Exec(lastInsertId, v.IpStr, v.Cidr)
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}
	p.Close()

	o.Commit()
	return rs, err
}

func UpdateMember(m *objs.Member, admin *objs.Member) (sql.Result, error) {

	o := orm.NewOrm()
	o.Begin()

	var query string

	// 사용자 정보 업데이트
	query = "update mbr_member set email = ?, name = ?, timezone = ? where member_id = ?"
	rs, err := o.Raw(query, m.Email, m.Name, m.Timezone, m.MemberId).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	// 사용자 권한 업데이트
	query = "update mbr_member set position = ? where member_id = ? and position < ?"
	rs, err = o.Raw(query, m.Position, m.MemberId, admin.Position).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	// 비밀번호 업데이트
	if len(m.Password) > 0 {
		query = "update mbr_password set password = ? where member_id = ?"
		rs, err = o.Raw(query, m.PasswordConfirm, m.MemberId).Exec()
		if err != nil {
			o.Rollback()
			return rs, err
		}

		query = "update mbr_member set failed_login_count = 0 where member_id = ?"
		rs, err = o.Raw(query, m.MemberId).Exec()
		if err != nil {
			o.Rollback()
			return rs, err
		}
	}

	// 접속허용 IP 삭제 후 업데이트
	o.Raw("delete from mbr_allowed_ip where member_id = ?", m.MemberId).Exec()
	p, err := o.Raw("insert into mbr_allowed_ip(member_id, ip, cidr) values(?, inet_aton(?), ?)").Prepare()
	for _, v := range m.AllowedIpList {
		_, err := p.Exec(m.MemberId, v.IpStr, v.Cidr)
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}
	p.Close()

	o.Commit()
	return rs, err
}

func AfterSignin(m *objs.Member) error {
	query := `
		update mbr_member
		set failed_login_count = 0, last_success_login = now(), login_count = login_count + 1, session_id = ?
		where member_id = ?
	`
	o := orm.NewOrm()
	_, err := o.Raw(query, m.SessionId, m.MemberId).Exec()
	return err
}

func LoginFailed(username string, lastFailedLogin bool) error {
	var s string

	if lastFailedLogin {
		s = ", last_failed_login = now()"
	}
	query := `
		update mbr_member
		set failed_login_count = failed_login_count + 1 %s
		where username = ?
	`
	query = fmt.Sprintf(query, s)
	o := orm.NewOrm()
	_, err := o.Raw(query, username).Exec()
	return err
}

func RemoveMember(memberId int, adminPosition int) (sql.Result, error) {
	query := "delete from mbr_member where member_id = ? and position < ?" // 삭제수행 주체보다 하위 권한만 삭제 가능

	o := orm.NewOrm()
	return o.Raw(query, memberId, adminPosition).Exec()
}

func GetMemberAcl(memberId int) ([]int, error) {
	query := "select asset_id from mbr_asset where member_id = ?"

	var list []int
	o := orm.NewOrm()
	_, err := o.Raw(query, memberId).QueryRows(&list)
	return list, err
}


// 사용자
func UpdateMemberAcl(memberId int, acl []string) (sql.Result, error) {
	o := orm.NewOrm()
	o.Begin()

	// 기존 자산정보 삭제
	query := "delete from mbr_asset where member_id = ?"
	rs, err := o.Raw(query,memberId).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	// 자산정보 업데이트(그룹 정보(type1=2)만 업데이트)
	query =`
		insert into mbr_asset(member_id, asset_id)
		select ?, asset_id
		from ast_asset
		where class = 1 and type1 = 2 and asset_id in (%s)
	`
	query = fmt.Sprintf(query, strings.Join(acl, ","))
	rs, err = o.Raw(query,memberId).Exec()
	if err != nil {
		o.Rollback()
		return rs, err
	}

	o.Commit()
	return rs, err
}
