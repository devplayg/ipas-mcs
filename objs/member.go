package objs

import (
	"errors"
	"github.com/devplayg/ipas-mcs/libs"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type Member struct {
	MemberId         int           `json:"member_id" form:"member_id"`
	Username         string         `json:"username" form:"username"`
	Password         string         `json:"-" form:"password"`
	PasswordConfirm  string         `json:"-" form:"password_confirm"`
	Salt             string         `json:"-"`
	Name             string         `json:"name" form:"name"`
	Position         uint           `json:"position"`
	Timezone         string         `json:"timezone" form:"timezone"`
	Location         *time.Location `json:"-"`
	SessionId        string         `json:"-"`
	FailedLoginCount uint           `json:"failed_login_count"`
	Email            string         `json:"email" form:"email"`
	AllowedIp        string         `json:"allowed_ip" form:"allowed_ip"`
	AllowedIpList    []IpCidr       `json:"-"`
	UserGroups       []int          `form:"user_groups"`
}

type IpCidr struct {
	IpStr string
	Cidr  int
}

func (m *Member) Validate() error {

	if len(m.Password) > 0 && len(m.PasswordConfirm) > 0 {
		//reUsername := regexp.MustCompile("^[[:alpha:]]{1}[[:word:]]{3,16}$")
		if !regexp.MustCompile(libs.UsernamePattern).MatchString(m.Username) {
			return errors.New("invalid username")
		}

		if !isValidPassword(m.Password) {
			return errors.New("invalid password")
		}
	}
	nameLength := utf8.RuneCountInString(m.Name)
	if !(nameLength >= 2 && nameLength <= 32) {
		return errors.New("invalid name's length(2-32)")
	}

	if !regexp.MustCompile(libs.EmailPattern).MatchString(m.Email) {
		return errors.New("invalid email")
	}

	return nil
}

func isValidPassword(str string) bool {
	hasLowerCase := false
	hasUpperCase := true
	hasNumber := false
	hasSpecialChar := false
	properLength := false

	specialChars := "!@#$%^&*()~_+`-=,.<>/?|"
	for _, s := range str {
		if unicode.IsLower(s) {
			hasLowerCase = true
			//log.Debugf("%#U lower", s)
		} else if unicode.IsUpper(s) {
			hasUpperCase = true
			//log.Debugf("%#U upper", s)
		} else if unicode.IsNumber(s) {
			hasNumber = true
			//log.Debugf("%#U number", s)
		} else if strings.ContainsRune(specialChars, s) {
			hasSpecialChar = true
			//log.Debugf("%#U symbol", s)
		} else {
			//log.Debugf("%#U nothing", s)
		}
	}

	if len(str) >= 8 && len(str) <= 16 {
		properLength = true
	}

	return hasLowerCase && hasUpperCase && hasNumber && hasSpecialChar && properLength
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
