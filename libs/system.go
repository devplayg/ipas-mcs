package libs

import (
	"fmt"
	"strings"
	"regexp"
)

func JoinInt(arr []int, delimiter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), delimiter), "[]")
}

func SplitString(s string, pattern string) []string {
	list := regexp.MustCompile(pattern).Split(strings.TrimSpace(s), -1)
	for i, v := range list {
		if v == "" {
			list = append(list[:i], list[i+1:]...)
		}
	}

	return list
}