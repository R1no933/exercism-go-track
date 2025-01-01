package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("input must be a letter between 'A' and 'Z'")
	}

	n := int(char - 'A')
	var rows []string
	for i := 0; i <= n; i++ {
		letter := string(byte('A' + i))
		spaces := strings.Repeat(" ", n-i)
		if i == 0 {
			rows = append(rows, spaces+letter+spaces)
		} else {
			innerSpaces := strings.Repeat(" ", 2*i-1)
			rows = append(rows, spaces+letter+innerSpaces+letter+spaces)
		}
	}
	for i := n - 1; i >= 0; i-- {
		rows = append(rows, rows[i])
	}
	return strings.Join(rows, "\n"), nil
}
