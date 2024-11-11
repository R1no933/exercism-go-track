package pangram

import "strings"

func IsPangram(input string) bool {
	inToLower := strings.ToLower(input)

	for ch := 'a'; ch <= 'z'; ch++ {
		if !strings.ContainsRune(inToLower, ch) {
			return false
		}
	}

	return true
}
