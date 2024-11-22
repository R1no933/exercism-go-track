package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	cntr := make(Frequency)
	reg := regexp.MustCompile(`\w+('\w+)?`)

	for _, wd := range reg.FindAllString(strings.ToLower(phrase), -1) {
		cntr[wd]++
	}

	return cntr
}
