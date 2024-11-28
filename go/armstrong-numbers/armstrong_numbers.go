package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	nStr := strconv.Itoa(n)
	lenN := len(nStr)
	armstrN := 0

	for _, ch := range nStr {
		chNmb, _ := strconv.Atoi(string(ch))
		armstrN += int(math.Pow(float64(chNmb), float64(lenN)))
	}

	if n != armstrN {
		return false
	}

	return true
}
