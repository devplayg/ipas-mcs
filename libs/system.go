package libs

import (
	"fmt"
	"strings"
	"regexp"
	"math/rand"
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_")

func GetRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
