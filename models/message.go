package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/ipas-mcs/objs"
	"time"
	"github.com/devplayg/ipas-server"
)

func GetMessage(filter objs.MessageFilter, member *objs.Member) ([]objs.Message, int64, error) {

	var where string
	var rows []objs.Message

	// 조건 설정
	args := make([]interface{}, 0)
	
	// 시간설정
	startDate, _ := time.ParseInLocation(ipasserver.DateDefault, filter.StartDate+":00", member.Location)
	endDate, _ := time.ParseInLocation(ipasserver.DateDefault, filter.EndDate+":59", member.Location)
	args = append(args, startDate.UTC().Format(ipasserver.DateDefault), endDate.UTC().Format(ipasserver.DateDefault), member.MemberId)

	if filter.Status > 0 {
		where += " and status = ?"
		args = append(args, filter.Status)
	}

	// 페이징 모드(고속/일반)
	if filter.FastPaging == "off" {
		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
	}

	// Set query
	query := `
		select %s message_id, date, status, receiver_id, sender_id, priority, category, message, url
		from log_message
		where date >= ? and date <= ? and receiver_id = ?  %s 
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

func MarkMessageAsRead(messageId int, member *objs.Member) error {
	query := "update log_message set status = 2 where message_id = ? and receiver_id = ?"
	args := []interface{}{messageId, member.MemberId}

	o := orm.NewOrm()
	_, err := o.Raw(query, args).Exec()
	return err
}

func MarkAllMessageAsRead(member *objs.Member) error {
	query := "update log_message set status = 2 where date >= date_add(now(), interval -7 day) and receiver_id = ?"
	o := orm.NewOrm()
	_, err := o.Raw(query, member.MemberId).Exec()
	spew.Dump(query)
	spew.Dump(member.MemberId)
	return err
}
