package objs

import (
	"time"
)

type Ipas struct {
	EquipId   string `json:"equip_id"`
	EquipType int    `json:"equip_type"`
	GroupId   int    `json:"group_id"`
	Speed     int    `json:"spped"`
	Snr       int    `json:"snr"`
	Usim      string `json:"usim"`
	//SpeedingCount int       `json:"speeding_count"`
	//ShockCount    int       `json:"shock_count"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
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

	Org       []int
	SubOrg    []int
	RiskLevel []int  `form:"risk_level[]"`
	Contents  string `form:"contents"`
	EquipId   string `form:"equip_id"`
}
