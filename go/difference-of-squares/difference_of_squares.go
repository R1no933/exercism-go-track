package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	res := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	res = sum * sum
	return res
}

func SumOfSquares(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res = res + i*i
	}

	return res
}

func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n) - SumOfSquares(n))))
}
