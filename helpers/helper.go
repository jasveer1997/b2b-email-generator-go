package helpers

import (
	"strconv"
)

func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func ParseStrToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}
	return int32(i)
}

func ContainsMatchingDomain(arr []string, val string) bool {
	for _, domain := range arr {
		if domain == val {
			return true
		}
	}
	return false
}
