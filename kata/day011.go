package kata

import (
	"regexp"
)

// Alphanumeric returns whether a string consists of only alphanumeric characters
func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	match, err := regexp.Match(`[_\W]`, []byte(str))
	if err != nil {
		return false
	}
	return !match
}
