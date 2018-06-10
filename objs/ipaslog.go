package objs

import (
	"time"
)

const (
	StartupEvent   = 1 // 시동
	ShockEvent     = 2 // 충격
	SpeedingEvent  = 3 // 과속
	ProximityEvent = 4 // 근접
)

type IpasLog struct {
	Ipas
	Date      time.Time `json:"date"`
	//OrgId     int       `json:"org_id"`
	//OrgName   string    `json:"org_name"`
	//GroupId   int       `json:"group_id"`
	//GroupName string    `json:"group_name"`
	EventType int       `json:"event_type"`
	SessionId string    `json:"session_id"`
	//EquipId   string    `json:"equip_id"`
	Targets   string    `json:"targets"`
	//Latitude  float32   `json:"latitude"`
	//Longitude float32   `json:"longitude"`
	//Speed     int       `json:"speed"`
	//Snr       int       `json:"snr"`
	//Usim      string    `json:"usim"`
	Distance  int       `json:"distance"`
	Ip        uint32    `json:"ip"`
	RecvDate  time.Time `json:"recv_date"`
	No        int64     `json:"no"`
	DateAgo   string    `json:"date_ago"`
}

type LocTrack struct {
	Date      time.Time `json:"date"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
}

// 시동,충격,과속,근접
//type Ipas struct {
//	OrgId     int     `json:"org_id"`
//	EquipId   string  `json:"equip_id"`
//	GroupId   int     `json:"group_id"`
//	EquipType int     `json:"equip_type"`
//	Speed     int     `json:"spped"`
//	Snr       int     `json:"snr"`
//	Usim      string  `json:"usim"`
//	Latitude  float32 `json:"latitude"`
//	Longitude float32 `json:"longitude"`
//	Created   time.Time `json:"created"`
//	Updated   time.Time `json:"updated"`
//	OrgName   string    `json:"org_name"`
//	GroupName string    `json:"group_name"`
//}



type IpasFilter struct {
	PagingFilter

	OrgId      []int  `form:"org_id"`
	GroupId    []int  `form:"group_id"`
	EventType  []int  `form:"event_type"`
	Contents   string `form:"contents"`
	EquipId    string `form:"equip_id"`
	TagPattern string `form:"tag_pattern"`
	StatsMode  bool   `form:"stats_mode"` // 통계모드
}
