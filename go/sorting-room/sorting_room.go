package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	result := 0

	switch tp := fnb.(type) {
	case FancyNumber:
		result, _ = strconv.Atoi(tp.Value())
	default:
		result = 0
	}

	return result
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	resStr := ""

	switch tp := i.(type) {
	case int:
		resStr = DescribeNumber(float64(tp))
	case float64:
		resStr = DescribeNumber(tp)
	case NumberBox:
		resStr = DescribeNumberBox(tp)
	case FancyNumberBox:
		resStr = DescribeFancyNumberBox(tp)
	default:
		resStr = "Return to sender"
	}

	return resStr
}
