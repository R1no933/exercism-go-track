package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var sb strings.Builder
	for i, n := 1, 1; i <= len(input); i++ {
		if i == len(input) || input[i] != input[i-1] {
			if n > 1 {
				sb.WriteString(strconv.Itoa(n))
				n = 1
			}
			sb.WriteByte(input[i-1])
		} else {
			n++
		}
	}
	return sb.String()
}

func RunLengthDecode(input string) string {
	var sb strings.Builder
	for i, n := 0, 0; i < len(input); i++ {
		if input[i] < 0x30 || input[i] > 0x39 {
			if n < i {
				n, _ = strconv.Atoi(input[n:i])
				sb.WriteString(strings.Repeat(string(input[i]), n))
			} else {
				sb.WriteByte(input[i])
			}
			n = i + 1
		}
	}
	return sb.String()
}
