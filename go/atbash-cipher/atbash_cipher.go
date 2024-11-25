package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var ans []string
	var prt string

	for _, ch := range s {
		switch {
		case unicode.IsDigit(ch):
			prt += string(ch)
		case unicode.IsUpper(ch):
			prt += string('A' + 'Z' - ch - ('A' - 'a'))
		case unicode.IsLower(ch):
			prt += string('a' + 'z' - ch)
		}

		if len(prt) == 5 {
			ans = append(ans, prt)
			prt = ""
		}
	}

	if len(prt) > 0 {
		ans = append(ans, prt)
	}

	return strings.Join(ans, " ")
}
