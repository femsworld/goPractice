package sprint

func Payout(amount int, denominations []int) []int {
	if amount < 0 || len(denominations) == 0 || !isPositiveDenominations(denominations) {
		return []int{}
	}
	for i := 0; i < len(denominations)-1; i++ {
		for j := 0; j < len(denominations)-i-1; j++ {
			if denominations[j] < denominations[j+1] {
				denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
			}
		}
	}

	result := []int{}

	for i := 0; i < len(denominations); i++ {
		for amount >= denominations[i] {
			result = append(result, denominations[i])
			amount -= denominations[i]
		}
	}

	if amount == 0 {
		return result
	}

	return []int{}
}

func isPositiveDenominations(denominations []int) bool {
	for _, denom := range denominations {
		if denom <= 0 {
			return false
		}
	}
	return true
}
