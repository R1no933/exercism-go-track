package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cowCount int
	message  string
}

func (cowErr *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", cowErr.cowCount, cowErr.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowCount int) (float64, error) {
	if fodder, err := fc.FodderAmount(cowCount); err == nil {
		if fatFactor, err := fc.FatteningFactor(); err == nil {
			return fodder * fatFactor / float64(cowCount), nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowCount int) (float64, error) {
	if cowCount <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cowCount)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowCount int) error {
	if cowCount < 0 {
		return &InvalidCowsError{cowCount: cowCount, message: "there are no negative cows"}
	}

	if cowCount == 0 {
		return &InvalidCowsError{cowCount: cowCount, message: "no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
