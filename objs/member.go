package objs

import (
	"errors"
	"unicode/utf8"
	//log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
	"unicode"
	"github.com/devplayg/ipas-mcs/libs"
	"net"
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

func (m Member) Validate() error {

	//reUsername := regexp.MustCompile("^[[:alpha:]]{1}[[:word:]]{3,16}$")
	if ! regexp.MustCompile(libs.UsernamePattern).MatchString(m.Username) {
		return errors.New("invalid username")
	}

	if !isValidPassword(m.Password) {
		return errors.New("invalid password")
	}

	nameLength := utf8.RuneCountInString(m.Name)
	if ! (nameLength >= 3 && nameLength <= 16) {
		return errors.New("invalid name")
	}

	if ! regexp.MustCompile(libs.EmailPattern).MatchString(m.Email) {
		return errors.New("invalid email")
	}

	m.AllowedIp = strings.TrimSpace(m.AllowedIp)
	list := regexp.MustCompile(`[\s|,]+`).Split( m.AllowedIp, -1)
	for _, s := range list {
		if len(s) > 0 {
			if strings.Index(s, "/") > -1 { // CIDR이 있으면
				_, _, err := net.ParseCIDR(s)
				if err != nil {
					return err
				}
			} else { // IP면
				ip := net.ParseIP(s)
				if ip == nil {
					return errors.New("invalid IP address: "+s)
				}
			}
		}
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
