package objs

import "time"

type Stats struct {
	Date      time.Time `json:"date"`
	AssetId   int       `json:"asset_id"`
	Item      string    `json:"item"`
	//ItemText  string    `json:"item_text"`
	Count     int       `json:"count"`
	Rank      int       `json:"rank"`
	OrgName   string    `json:"org_name"`
	GroupName string    `json:"group_name"`
}
