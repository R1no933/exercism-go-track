package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float32

	switch {
	case balance < 0:
		rate = 3.213
	case balance >= 0 && balance < 1000:
		rate = 0.5
	case balance >= 1000 && balance < 5000:
		rate = 1.621
	default:
		rate = 2.475
	}

	return rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
	return balance * float64(rate) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	if balance >= targetBalance {
		years = 0
	}

	newBalance := AnnualBalanceUpdate(balance)
	for i := 1; newBalance < targetBalance; i++ {
		newBalance = newBalance + Interest(newBalance)
		years = i + 1
	}

	return years
}
