package libs

// https://github.com/google/re2/wiki/Syntax

const (
	EmailPattern = `^[a-zA-Z0-9.!#$%&'*+\/=?^_`+"`"+`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
)
