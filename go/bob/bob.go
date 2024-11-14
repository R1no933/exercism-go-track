// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

func (r Remark) empty() bool {
	return r.remark == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isShout() bool {
	hasLetter := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpper := strings.ToUpper(r.remark) == r.remark

	return hasLetter && isUpper
}

func (r Remark) isExasper() bool {
	return r.isQuestion() && r.isShout()
}

func newRemark(r string) Remark {
	return Remark{strings.TrimSpace(r)}
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.empty():
		return "Fine. Be that way!"
	case r.isExasper():
		return "Calm down, I know what I'm doing!"
	case r.isShout():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
