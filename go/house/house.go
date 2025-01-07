package house

import (
	"strings"
)

var lines = []string{
	"house that Jack built",
	"malt\nthat lay in",
	"rat\nthat ate",
	"cat\nthat killed",
	"dog\nthat worried",
	"cow with the crumpled horn\nthat tossed",
	"maiden all forlorn\nthat milked",
	"man all tattered and torn\nthat kissed",
	"priest all shaven and shorn\nthat married",
	"rooster that crowed in the morn\nthat woke",
	"farmer sowing his corn\nthat kept",
	"horse and the hound and the horn\nthat belonged to",
}

func Verse(v int) string {
	verse := "This is the " + lines[v-1]
	for i := v - 2; i >= 0; i-- {
		verse += " the " + lines[i]
	}
	return verse + "."
}
func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}
