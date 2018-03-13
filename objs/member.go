package objs

import (
	"errors"
	"github.com/devplayg/ipas-mcs/libs"
	"unicode/utf8"
	"github.com/go-ozzo/ozzo-validation"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
	"unicode"
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

func (a Member) Validate() error {

	if !isValidPassword(a.Password) {
		return errors.New("invalid password")
	}
	log.Debugf("Name count: %d ", utf8.RuneCountInString(a.Name))

	return validation.ValidateStruct(&a,
		validation.Field(
			&a.Username,
			validation.Required,
			validation.Match(regexp.MustCompile("^[[:alpha:]]{1}[[:word:]]{3,16}$")),
		),
		validation.Field(
			&a.Password,
			validation.Required,
			validation.Length(9, 16),
		),
		validation.Field(
			&a.Name,
			validation.Required,
			validation.RuneLength(2, 16),
		),
		validation.Field(
			&a.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(libs.EmailPattern)),
			validation.Length(8, 254),
		),
	)
}

func isValidPassword(str string) bool {
	hasLowerCase := false
	hasUpperCase := true
	hasNumber := false
	hasSpecialChar := false

	specialChars := "!@#$%^&*()~_+`-=,.<>/?|"
	for _, s := range str {
		if unicode.IsLower(s) {

			hasLowerCase = true
			log.Debugf("%#U lower", s)
		} else if unicode.IsUpper(s) {
			hasUpperCase = true
			log.Debugf("%#U upper", s)
		} else if unicode.IsNumber(s) {
			hasNumber = true
			log.Debugf("%#U number", s)
		} else if strings.ContainsRune(specialChars, s) {
			hasSpecialChar = true
			log.Debugf("%#U symbol", s)
		} else {
			log.Debugf("%#U nothing", s)
		}
	}

	return hasLowerCase && hasUpperCase && hasNumber && hasSpecialChar
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
