package objs

import (
	"time"
)

type Member struct {
	MemberId         int            `json:"member_id"`
	Username         string         `json:"username" form:"username"`
	Password         string         `json:"-" form:"password"`
	PasswordConfirm  string         `json:"-" form:"password_confirm"`
	Salt             string         `json:"-"`
	Name             string         `json:"name" form:"name"`
	Position         int            `json:"position"`
	Timezone         string         `json:"timezone" form:"timezone"`
	Location         *time.Location `json:"-"`
	SessionId        string         `json:"-"`
	FailedLoginCount int            `json:"failed_login_count"`
	Email            string         `json:"email" form:"email"`
	AllowedIp        string         `json:"allowed_ip" form:"allowed_ip"`
	UserGroups       []int          `form:"user_groups"`
	//OldPassword        string `json:"-"`
	//NewPassword        string `json:"-"`
	//NewPasswordConfirm string `json:"-"`
	//EncPassword        string `json:"-"`
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
