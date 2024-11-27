package sieve

func Sieve(limit int) []int {
	sv := make([]bool, limit+1)
	pms := make([]int, limit/2)
	idx := 0
	for i := 2; i <= limit; i++ {
		if !sv[i] {
			pms[idx] = i
			idx++
			for j := i + i; j <= limit; j += i {
				sv[j] = true
			}
		}
	}
	return pms[:idx]
}
