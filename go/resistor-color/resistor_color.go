package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	switch {
	case color == "black":
		return 0
	case color == "brown":
		return 1
	case color == "red":
		return 2
	case color == "orange":
		return 3
	case color == "yellow":
		return 4
	case color == "green":
		return 5
	case color == "blue":
		return 6
	case color == "violet":
		return 7
	case color == "grey":
		return 8
	case color == "white":
		return 9
	default:
		return 000
	}
}
