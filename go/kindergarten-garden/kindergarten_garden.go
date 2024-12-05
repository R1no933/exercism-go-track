package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden map of string to []string represents children gardening project
type Garden map[string][]string

// Plants list of allowed plants
var PlantsList = []string{"grass", "clover", "radishes", "violets"}

// NewGarden returns new Garden representation
//
// representation of diagram is empty line followed by garden representation
// diagram represents plants of children(children in alphabetical order)
// check for correct formed input
// - diagram must have correct form
// - diagram must correspond to right number of childen
// - values of diagram must be valid
// - children may not have same names
func NewGarden(diagram string, children []string) (*Garden, error) {
	// generate hash table of plants
	plants := map[rune]string{}
	for _, p := range PlantsList {
		plants[rune(p[0])-'a'+'A'] = p
	}
	// check formal diagram size
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 || rows[0] != "" {
		return nil, errors.New("diagram must have two rows, on three lines, first empty")
	}
	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("diagram rows 1 + 2 must have same length")
	}
	if len(rows[1]) != len(children)*2 {
		return nil, errors.New("diagram row mus be length of children*2")
	}
	g := Garden{}
	// sort children
	cnames := append([]string{}, children...)
	sort.Strings(cnames)
	// initial garden with placeholders for children plants
	for _, cn := range cnames {
		g[cn] = make([]string, 0, 4)
	}
	if len(g) != len(children) {
		return nil, errors.New("names of children are not unique")
	}
	for _, row := range rows[1:] {
		for cnx, cn := range cnames {
			p1, p2 := plants[rune(row[cnx*2])], plants[rune(row[cnx*2+1])]
			if p1 == "" || p2 == "" {
				return nil, errors.New("diagram contains not a plant")
			}
			g[cn] = append(g[cn], p1, p2)
		}
	}
	return &g, nil
}

// Plants is a Getter for Garden, returning plants per child
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
