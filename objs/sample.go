package objs

import "time"

type SampleLog struct {
	No        uint64    `json:"no"`
	Date      time.Time `json:"date"`
	RecvDate  time.Time `json:"recv_date"`
	Org       int       `json:"org"`
	SubOrg    int       `json:"sub_org"`
	Guid      string    `json:"guid"`
	RiskLevel int       `json:"risk_level"`
	Contents  string    `json:"contents"`
}

type SampleFilter struct {
	PagingFilter

	Org       []int
	SubOrg    []int
	RiskLevel []int  `form:"risk_level[]"`
	Contents  string `form:"contents"`
	Guid      string `form:"guid"`
}
