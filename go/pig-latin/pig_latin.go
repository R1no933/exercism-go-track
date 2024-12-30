package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	rl1 := regexp.MustCompile(`(?i)^(a|e|i|o|u|xr|yt)`)
	rl2rl3 := regexp.MustCompile(`(?i)^([b-df-hj-np-tv-xz]*qu|[b-df-hj-np-tv-z]+)(.*)$`)
	rl4 := regexp.MustCompile(`(?i)^([b-df-hj-np-tv-xz]+)(y.*)$`)
	apply := func(m []string) string {
		return m[2] + m[1] + "ay"
	}
	words := strings.Fields(sentence)
	for i, word := range words {
		if rl1.MatchString(word) {
			words[i] += "ay"
		} else if m := rl4.FindStringSubmatch(word); m != nil {
			words[i] = apply(m)
		} else {
			words[i] = apply(rl2rl3.FindStringSubmatch(word))
		}
	}

	return strings.Join(words, " ")
}
