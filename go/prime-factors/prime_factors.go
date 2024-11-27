package prime

func Factors(n int64) []int64 {
	nmbrs := make([]int64, 0)

	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			nmbrs = append(nmbrs, i)
			n /= i
		}
	}
	return nmbrs
}
