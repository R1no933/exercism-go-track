package robotname

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Robot struct {
	name string
}

const (
	nameLimit = 26 * 26 * 1000
)

var (
	alphabet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	idx      = 0
	pool     = generate()
)

func generate() []string {
	names := make([]string, nameLimit)
	pos := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 1000; k++ {
				name := fmt.Sprintf("%c%c%03d", alphabet[i], alphabet[j], k)
				names[pos] = name
				pos++
			}
		}
	}
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}
func (r *Robot) Name() (string, error) {
	if r != nil && r.name != "" {
		return r.name, nil
	}
	if idx >= nameLimit {
		return "", errors.New("names exhausted")
	}
	r.name = pool[idx]
	idx++
	return r.name, nil
}
func (r *Robot) Reset() {
	r.name = ""
}
