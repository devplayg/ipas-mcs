package objs

import (
	"time"
)

type Member struct {
	MemberId           int
	Username           string
	//Password           string `json:"-"`
	//PasswordConfirm    string `json:"-"`
	//OldPassword        string `json:"-"`
	//NewPassword        string `json:"-"`
	//NewPasswordConfirm string `json:"-"`
	//EncPassword        string `json:"-"`
	//Email              string
	//Salt               string `json:"-"`
	Name               string
	//Status             int `json:"-"`
	Position           int
	//FailedLoginCount   int
	Timezone           string `json:"-"`
	//AllowedIp          string
	//Usergroups         []int `json:"-"`
	aaa string
	Location           *time.Location
	SessionId          string `json:"-"`
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