package isbn

func getDigit(ch rune) (digit int, ok bool) {
	if ch < '0' || ch > '9' {
		return digit, ok
	}
	return int(ch - '0'), true
}

func IsValidISBN(input string) bool {
	i := 10
	sum := 0

	for _, ch := range input {
		if digit, ok := getDigit(ch); ok {
			sum += (digit * i)
			i--
		} else if i == 1 && ch == 'X' {
			sum += 10
			i--
		} else if ch != '-' {
			return false
		}
	}
	return i == 0 && sum%11 == 0
}
