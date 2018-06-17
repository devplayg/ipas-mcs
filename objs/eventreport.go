package objs

type ReportFilter struct {
	//Date         time.Time `form:"date"`
	Date      string `form:"date"`
	StartDate string
	EndDate   string
	OrgId     int    `form:"org_id"`
	GroupId   int    `form:"group_id"`
	EquipId   string `form:"equip_id"`
	PastDays  int    `form:"past_days"`
}
