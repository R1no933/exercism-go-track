package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must ne greater than 1")
	}

	nth := 0
	cnt := 0

	for {
		cnt++
		if isPrimeNumber(cnt) {
			nth++
		}

		if nth == n {
			return cnt, nil
		}
	}
}

func isPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
