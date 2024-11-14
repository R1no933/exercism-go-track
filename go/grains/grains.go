package grains

import (
	"errors"
	"math"
	"math/big"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("square number must be 1 through 64")
	}
	return uint64(math.Pow(2, float64(number)-1)), nil
}

func Total() uint64 {
	nmb, exp := big.NewInt(2), big.NewInt(64)
	return nmb.Exp(nmb, exp, nil).Sub(nmb, big.NewInt(1)).Uint64()
}
