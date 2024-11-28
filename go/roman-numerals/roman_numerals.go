package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("Input number must be less than 3999")
	}

	var convertResult strings.Builder
	romanSymb := []struct {
		val  int
		symb string
	}{{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, smb := range romanSymb {
		for input >= smb.val {
			convertResult.WriteString(smb.symb)
			input -= smb.val
		}
	}

	return convertResult.String(), nil
}
