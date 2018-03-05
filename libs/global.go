package libs


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

// Config
var config = make(map[string]map[string]configValue)
type configValue struct {
	ValueS string
	ValueN int
}

// Multi-language
type langType struct {
	Lang, Name string
}
var langTypes []*langType // Languages are supported.

