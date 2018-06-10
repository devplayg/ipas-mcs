package objs

import "time"

type ReportFilter struct {
	Date      time.Time `form:"date"`
	StartDate string
	EndDate   string
	OrgId     int       `form:"org_id"`
	GroupId   int       `form:"group_id"`
	EquipId   string    `form:"equip_id"`
	SinceDays int       `form:"since_days"`
}
