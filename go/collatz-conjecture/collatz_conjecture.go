package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("N must be more than 0")
	}

	cnt := 0
	for i := 0; n != 1; i++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		cnt++
	}

	return cnt, nil
}
