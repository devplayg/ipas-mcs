package objs

import "time"

type Stats struct {
	Date      time.Time `json:"date"`
	AssetId   int       `json:"asset_id"`
	Item      string    `json:"item"`
	Count     int       `json:"count"`
	Rank      int       `json:"rank"`
	OrgName   string    `json:"org_name"`
	GroupName string    `json:"group_name"`
}

type StatsFilter struct {
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	Top       int    `form:"top"`
	StatsType string
	AssetType string
	OrgId     int
	GroupId   int
}

type TagCount struct {
	EquipType int `json:"equip_type"`
	Count     int `json:"count"`
}
