package bottlesong

import (
	"fmt"
	"strings"
)

var song = `%s green bottle%s hanging on the wall,
%s green bottle%s hanging on the wall,
And if one green bottle should accidentally fall,
There'll be %s green bottle%s hanging on the wall.

`
var bottlesCount = []string{
	"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
}

func Recite(startBottles, takeDown int) []string {
	s := ""
	for i := startBottles; i > startBottles-takeDown; i-- {
		apstphe := "s"
		if i == 1 {
			apstphe = ""
		}
		lstApstphe := "s"
		if i-1 == 1 {
			lstApstphe = ""
		}
		s += fmt.Sprintf(song,
			bottlesCount[i],
			apstphe,
			bottlesCount[i],
			apstphe,
			strings.ToLower(bottlesCount[i-1]),
			lstApstphe,
		)
	}
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}
