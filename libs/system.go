package libs

import (
	"fmt"
	"strings"
)

func JoinInt(arr []int, delimiter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), delimiter), "[]")
}
