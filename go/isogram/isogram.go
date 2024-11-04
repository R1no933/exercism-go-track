package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)

	for _, ch := range lowerWord {
		if !unicode.IsLetter(ch) {
			continue
		}

		if strings.Count(lowerWord, string(ch)) > 1 {
			return false
		}
	}

	return true
}
