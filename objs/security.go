package objs

import "time"

type SecurityLogFilter struct {
	PagingFilter
	Category []string `form:"category"`
}

type SecurityLog struct {
	Date     time.Time `json:"date"`
	AuditId  int `json:"audit_id"`
	MemberId int `json:"member_id"`
	Category string `json:"category"`
	IP       uint32 `json:"ip"`
	Message  string `json:"message"`
	Detail   string `json:"detail"`
}
