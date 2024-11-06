package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	sum := 0
	isSecondDigit := len(id)%2 == 0
	for _, r := range id {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if isSecondDigit {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
		isSecondDigit = !isSecondDigit
	}
	return sum%10 == 0
}
