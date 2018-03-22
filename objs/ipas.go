package objs

import (
	"time"
)

type Ipas struct {
	EquipId       string    `json:"equip_id"`
	GroupId       int       `json:"group_id"`
	Type          int       `json:"type"`
	SpeedingCount int       `json:"speeding_count"`
	ShockCount    int       `json:"shock_count"`
	Srn           string    `json:"srn"`
	Contact       string    `json:"contact"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}

type IpasLog struct {
	Date          time.Time `json:"date"`
	EquipId       string    `json:"equip_id"`
	Target        string    `json:"target"`
	SpeedingCount int       `json:"speeding_count"`
	ShockCount    int       `json:"shock_count"`
	Latitude      float32   `json:"latitude"`
	Longitude     float32   `json:"longitude"`
	WarningDist   int       `json:"warning_dist"`
	CautionDist   int       `json:"caution_dist"`
	V2vDist       int       `json:"v2v_dist"`
	CollisionThr  int       `json:"collision_thr"`
	ShockThr      int       `json:"shock_thr"`
	SpeedThr      int       `json:"speed_thr"`
	Rdate         time.Time `json:"rdate"`

	//No        uint64    `json:"no"`
	//Date      time.Time `json:"date"`
	//RecvDate  time.Time `json:"recv_date"`
	//Org       int       `json:"org"`
	//SubOrg    int       `json:"sub_org"`
	//Guid      string    `json:"guid"`
	//RiskLevel int       `json:"risk_level"`
	//Contents  string    `json:"contents"`
}

type IpasFilter struct {
	PagingFilter

	//Org       []int
	//SubOrg    []int
	//RiskLevel []int  `form:"risk_level[]"`
	//Contents  string `form:"contents"`
	//Guid      string `form:"guid"`
}
