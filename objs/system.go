package objs

const (
	User = 1 << iota
	UnknownLeve2
	UnknownLeve3
	UnknownLeve4
	UnknownLeve5
	UnknownLeve6
	UnknownLeve7
	UnknownLeve8
	Observer      // 256
	Administrator // 512
	Superman      // 1024
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

const (
	DefaultDateFormat = "2006-01-02 15:04:05"
)

// System configuration
type SysConfig struct {
	Section, Keyword, ValueS string
	ValueN                   int
}

var SysConfigMap = make(map[string]map[string]MultiValue)

type MultiValue struct {
	ValueS string
	ValueN int
}

type AuditMsg struct {
	MemberId int
	Category string
	IP       string
	Message  interface{}
	Detail   interface{}
}

type PagingFilter struct {
	StartDate  string `form:"startDate"`
	EndDate    string `form:"endDate"`
	FastPaging string `form:"fastPaging"`
	FoundRows  string

	// Paging
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`

	// Sort
	Order  string `form:"order"`
	Sort   string `form:"sort"`
}
