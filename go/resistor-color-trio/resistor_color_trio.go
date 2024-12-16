package resistorcolortrio

import (
	"fmt"
	"math"
)

var colorLookup = map[string]float64{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Label(colors []string) string {
	coeff := colorLookup[colors[0]]*10 + colorLookup[colors[1]]
	exp := colorLookup[colors[2]]
	vl := coeff * math.Pow(10, exp)
	units := "ohms"
	if vl > 999_999_999 {
		vl /= 1_000_000_000
		units = "gigaohms"
	} else if vl > 999_999 {
		vl /= 1_000_000
		units = "megaohms"
	} else if vl > 999 {
		vl /= 1_000
		units = "kiloohms"
	}
	return fmt.Sprintf("%.0f %s", vl, units)
}
