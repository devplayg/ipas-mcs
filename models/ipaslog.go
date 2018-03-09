package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/devplayg/ipas-mcs/libs"
	"github.com/devplayg/ipas-mcs/objs"
)

func GetIpaslog(filter *objs.IpasFilter) ([]objs.IpasLog, int64, error) {
	var where string
	var rows []objs.IpasLog

	args := make([]interface{}, 0, 20)

	if len(filter.Orgs) > 0 {
		where += fmt.Sprintf(" and org in (%s)", libs.JoinInt(filter.Orgs, ","))
	}

	// Set query
	query := `
		SELECT %s *
		from log_ipas
		where date >= ? and date <= ? %s order by %s %s limit ?, ?
	`
	query = fmt.Sprintf(query, filter.FoundRows, where, filter.Sort, filter.Order)

	args = append(args, filter.StartDate + ":00", filter.EndDate+":59", filter.Offset, filter.Limit)

	o := orm.NewOrm()
	total, _ := o.Raw(query, args).QueryRows(&rows)
	//	dbResult := DbResult{}
	//	isMatch := RegexFoundRows.MatchString(query)
	//	if isMatch {
	//		o.Raw("select FOUND_ROWS() total").QueryRow(&dbResult)
	//		total = dbResult.Total
	//
	//	}
	return rows, total, nil

}

//func GetIpasLog(filter *FilterFiletranslog) ([]objs.IpasLog, int64, error) {
//	var where string
//	args := make([]interface{}, 0, 20)
//
//	// Fast paging
//	if filter.FastPaging != "on" {
//		filter.FoundRows = "SQL_CALC_FOUND_ROWS"
//	}
//
//	// Time
//	args = append(args, filter.StartDate+":00")
//	args = append(args, filter.EndDate+":59")
//
//	// Network
//	if filter.Networks != nil {
//		where += fmt.Sprintf(" and network_id in (%s)", StringJoin(filter.Networks, ","))
//	} else {
//		if filter.Groups != nil {
//			where += fmt.Sprintf(" and group_id in (%s)", StringJoin(filter.Groups, ","))
//		} else {
//			if filter.Sensors != nil {
//				where += fmt.Sprintf(" and sensor_id in (%s)", StringJoin(filter.Sensors, ","))
//			} else {
//				// N/A
//			}
//		}
//
//	}
//	// Src IP
//	if len(filter.SrcIpCidr) > 0 {
//		if strings.Contains(filter.SrcIpCidr, "/") { // CIDR
//			ip_min, ip_max, err := libs.GetCidrMinMax(filter.SrcIpCidr)
//			CheckError(err)
//			if err == nil {
//				where += " and src_ip between ? and ?"
//				args = append(args, libs.IpToInt32(ip_min), libs.IpToInt32(ip_max))
//			}
//		} else { // Single IP
//			ip := net.ParseIP(filter.SrcIpCidr)
//			if ip != nil {
//				where += " and src_ip = ?"
//				args = append(args, libs.IpToInt32(ip))
//			}
//		}
//	}
//
//	// Src port
//	if len(filter.SrcPortStart) > 0 && len(filter.SrcPortEnd) > 0 {
//		where += " and src_port between ? and ?"
//		port_from, _ := strconv.Atoi(filter.SrcPortStart)
//		port_to, _ := strconv.Atoi(filter.SrcPortEnd)
//		args = append(args, port_from, port_to)
//	} else {
//		if len(filter.SrcPortStart) > 0 {
//			where += " and src_port >= ?"
//			port, _ := strconv.Atoi(filter.SrcPortStart)
//			args = append(args, port)
//		} else if len(filter.SrcPortEnd) > 0 {
//			where += " and src_port <= ?"
//			port, _ := strconv.Atoi(filter.SrcPortEnd)
//			args = append(args, port)
//		}
//	}
//
//	// Dst IP
//	if len(filter.DstIpCidr) > 0 {
//		if strings.Contains(filter.DstIpCidr, "/") { // CIDR
//			ip_min, ip_max, err := libs.GetCidrMinMax(filter.DstIpCidr)
//			CheckError(err)
//			if err == nil {
//				where += " and Dst_ip between ? and ?"
//				args = append(args, libs.IpToInt32(ip_min), libs.IpToInt32(ip_max))
//			}
//		} else { // Single IP
//			ip := net.ParseIP(filter.DstIpCidr)
//			if ip != nil {
//				where += " and Dst_ip = ?"
//				args = append(args, libs.IpToInt32(ip))
//			}
//		}
//	}
//
//	// Dst port
//	if len(filter.DstPortStart) > 0 && len(filter.DstPortEnd) > 0 {
//		where += " and Dst_port between ? and ?"
//		port_from, _ := strconv.Atoi(filter.DstPortStart)
//		port_to, _ := strconv.Atoi(filter.DstPortEnd)
//		args = append(args, port_from, port_to)
//	} else {
//		if len(filter.DstPortStart) > 0 {
//			where += " and Dst_port >= ?"
//			port, _ := strconv.Atoi(filter.DstPortStart)
//			args = append(args, port)
//		} else if len(filter.DstPortEnd) > 0 {
//			where += " and Dst_port <= ?"
//			port, _ := strconv.Atoi(filter.DstPortEnd)
//			args = append(args, port)
//		}
//	}
//
//	// Paging
//	args = append(args, filter.Offset)
//	args = append(args, filter.Limit)
//
//	// Set query
//	query := `
//        SELECT %s *
//        from log_event_filetrans
//        where rdate >= ? and rdate <= ? %s order by %s %s limit ?, ?
//    `
//	query = fmt.Sprintf(query, filter.FoundRows, where, filter.Sort, filter.Order)
//
//	o := orm.NewOrm()
//	rows := []RsFiletransLog{}
//	total, err := o.Raw(query, args).QueryRows(&rows)
//	dbResult := DbResult{}
//	isMatch := RegexFoundRows.MatchString(query)
//	if isMatch {
//		o.Raw("select FOUND_ROWS() total").QueryRow(&dbResult)
//		total = dbResult.Total
//
//	}
//
//	return &rows, total, err
//}
//
