package objs

import (
	"time"
)

type Member struct {
	MemberId int    `json:"member_id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Salt string `json:"-"`
	Name string `json:"name"`
	Position int `json:"position"`
	Timezone string `json:"-"`
	Location  *time.Location `json:"-"`
	SessionId string         `json:"-"`
	FailedLoginCount   int `json:"failed_login_count"`
	//PasswordConfirm    string `json:"-"`
	//OldPassword        string `json:"-"`
	//NewPassword        string `json:"-"`
	//NewPasswordConfirm string `json:"-"`
	//EncPassword        string `json:"-"`
	//Email              string
	//Status             int `json:"-"`
	//AllowedIp          string
	//Usergroups         []int `json:"-"`
}

//
//type MemberConfig struct {
//	MemberId int
//	Keyword  string
//	ValueS   string
//	ValueN   int
//}
//
//type MemberAsset struct {
//	MemberId int
//	Class    int
//	Assets   []int
//}
//
//type IpCidr struct {
//	Ip   string
//	Cidr int
//}
//
//type Password struct {
//	Password string
//	Salt     string
//}
