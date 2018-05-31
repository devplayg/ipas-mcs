package objs

import "time"

type Stats struct {
	Date      time.Time `json:"-"`
	AssetId   int       `json:"asset_id"`
	OrgId     int       `json:"org_id"`
	GroupId   int       `json:"group_id"`
	Item      string    `json:"item"`
	Count     int       `json:"count"`
	Rank      int       `json:"rank"`
	OrgName   string    `json:"org_name"`
	GroupName string    `json:"group_name"`

	// Timeline
	StartupCount   int `json:"startup_count"`
	ShockCount     int `json:"shock_count"`
	SpeedingCount  int `json:"speeding_count"`
	ProximityCount int `json:"proximity_count"`
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

type StatsByEventType struct {
	Date      time.Time `json:"date"`
	EventType int       `json:"event_type"`
	Count     int       `json:"count"`
}
