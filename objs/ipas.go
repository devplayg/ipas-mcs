package objs

import (
	"time"
)

type IpasLog struct {
	No        uint64    `json:"no"`
	Date      time.Time `json:"date"`
	RecvDate  time.Time `json:"recv_date"`
	Org       int       `json:"org"`
	SubOrg    int       `json:"sub_org"`
	Guid      string    `json:"guid"`
	RiskLevel int       `json:"risk_level"`
	Contents  string    `json:"contents"`
}

type IpasFilter struct {
	CommonFilter

	Orgs       []int
	SubOrgs    []int
	RiskLevels []int
	Contents   string
}
