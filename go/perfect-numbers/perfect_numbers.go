package perfect

import "errors"

// Define the Classification type here.

type Classification int

const (
	ClassificationPerfect   Classification = iota
	ClassificationAbundant  Classification = iota
	ClassificationDeficient Classification = iota
)

var ErrOnlyPositive = errors.New("invalid")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassificationDeficient, ErrOnlyPositive
	}
	sum := int64(0)
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			sum += i
		}
	}
	if sum > n {
		return ClassificationAbundant, nil
	} else if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
