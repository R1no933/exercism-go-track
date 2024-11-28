package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	firstColor := colors[0]
	secondColor := colors[1]
	result := ""
	switch firstColor {
	case "black":
		result += "0"
	case "brown":
		result += "1"
	case "red":
		result += "2"
	case "orange":
		result += "3"
	case "yellow":
		result += "4"
	case "green":
		result += "5"
	case "blue":
		result += "6"
	case "violet":
		result += "7"
	case "grey":
		result += "8"
	case "white":
		result += "9"
	}

	switch secondColor {
	case "black":
		result += "0"
	case "brown":
		result += "1"
	case "red":
		result += "2"
	case "orange":
		result += "3"
	case "yellow":
		result += "4"
	case "green":
		result += "5"
	case "blue":
		result += "6"
	case "violet":
		result += "7"
	case "grey":
		result += "8"
	case "white":
		result += "9"
	}

	intRes, _ := strconv.Atoi(result)

	return intRes
}
