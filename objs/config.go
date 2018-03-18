package objs

type Configuration struct {
	// 시스템
	DataRetentionDays      int    `form:"data_retention_days"`
	UseNameCard            string `form:"use_namecard"`

	// 로그인
	MaxFailedLoginAttempts int    `form:"max_failed_login_attempts"`
	AllowMultipleLogin     string `form:"allow_multiple_login"`
}
