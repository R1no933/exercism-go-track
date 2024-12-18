// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var sb strings.Builder

	for _, ch := range strings.FieldsFunc(s, Splitter) {
		rns := []rune(ch)
		sb.WriteRune(rns[0])
	}
	return strings.ToUpper(sb.String())
}

func Splitter(r rune) bool {
	return r == ' ' || r == '-' || r == '_'
}
