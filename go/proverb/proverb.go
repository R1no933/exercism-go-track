// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const proverbAll = "For want of a %s the %s was lost."
const proverbEnd = "And all for the want of a %s."

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))

	if len(rhyme) == 1 {
		proverb[0] += fmt.Sprintf(proverbEnd, rhyme[0])
	} else {
		for i := 0; i < len(rhyme); i++ {
			if i < len(rhyme)-1 {
				proverb[i] += fmt.Sprintf(proverbAll, rhyme[i], rhyme[i+1])
			} else {
				proverb[i] += fmt.Sprintf(proverbEnd, rhyme[0])
			}
		}
	}

	return proverb
}
