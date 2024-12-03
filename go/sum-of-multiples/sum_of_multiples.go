package summultiples

func SumMultiples(limit int, divisors ...int) int {
	magicScore := 0
	for i := 1; i < limit; i++ {
		for _, x := range divisors {
			if x != 0 && i%x == 0 {
				magicScore += i
				break
			}
		}
	}
	return magicScore
}
