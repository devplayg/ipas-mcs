package objs

type Configuration struct {
	// 시스템
	DataRetentionDays int `form:"data_retention_days"`
	// 로그인
	MaxFailedLoginAttempts int `form:"max_failed_login_attempts"`
	LoginFailureBlockTime  int `form:"login_failure_block_time"`
}
