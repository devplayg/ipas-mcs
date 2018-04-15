package objs

import (
	"time"
)

// 시동,충격,과속,근접
//select org_id, equip_id, group_id, equip_type, latitude, longitude, speed, snr, usim
//from ast_ipas
type Ipas struct {
	OrgId     int     `json:"org_id"`
	EquipId   string  `json:"equip_id"`
	GroupId   int     `json:"group_id"`
	EquipType int     `json:"equip_type"`
	Speed     int     `json:"spped"`
	Snr       int     `json:"snr"`
	Usim      string  `json:"usim"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	//SpeedingCount int       `json:"speeding_count"`
	//ShockCount    int       `json:"shock_count"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	OrgName   string    `json:"org_name"`
	GroupName string    `json:"group_name"`
}

type IpasLog struct {
	Date      time.Time `json:"date"`
	OrgId     int       `json:"org_id"`
	GroupId   int       `json:"group_id"`
	EventType int       `json:"event_type"`
	SessionId string    `json:"session_id"`
	EquipId   string    `json:"equip_id"`
	Targets   string    `json:"targets"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Speed     int       `json:"speed"`
	Snr       int       `json:"snr"`
	Usim      string    `json:"usim"`
	Distance  int       `json:"distance"`
	Ip        uint32    `json:"ip"`
	RecvDate  time.Time `json:"recv_date"`
}

type IpasFilter struct {
	PagingFilter

	OrgId     []int
	GroupId   []int
	EventType []int  `form:"event_type[]"`
	Contents  string `form:"contents"`
	EquipId   string `form:"equip_id"`
}
