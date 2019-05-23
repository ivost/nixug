package models

import (
	"strconv"
	"strings"
)

func safeStr(s string) string {
	return strings.Trim(s, " \t\n")
}

func safeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
