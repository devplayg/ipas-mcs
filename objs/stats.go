package objs

import "time"

type Stats struct {
	Date      time.Time `json:"date"`
	AssetId   int       `json:"asset_id"`
	OrgId     int       `json:"org_id"`
	GroupId   int       `json:"group_id"`
	Item      string    `json:"item"`
	Count     int       `json:"count"`
	Rank      int       `json:"rank"`
	OrgName   string    `json:"org_name"`
	GroupName string    `json:"group_name"`
	EquipId   string    `json:"equip_id"`
	Data      string    `json:"data"`

	// Timeline
	StartupCount   int `json:"startup_count"`
	ShockCount     int `json:"shock_count"`
	SpeedingCount  int `json:"speeding_count"`
	ProximityCount int `json:"proximity_count"`

	Uptime  int       `json:"uptime"`
}

type StatsFilter struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_Date"`
	Top       int    `form:"top"`
	StatsType string
	AssetType string
	OrgId     int
	GroupId   int
	EquipIp   string `json:"equip_id"`
}

type TagCount struct {
	EquipType int `json:"equip_type"`
	Count     int `json:"count"`
}

type StatsByEventType struct {
	Date      time.Time `json:"date"`
	EventType int       `json:"event_type"`
	Count     int       `json:"count"`
}
