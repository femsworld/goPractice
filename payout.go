package sprint

func Payout(amount int, denominations []int) []int {
	result := []int{}

	for i := len(denominations) - 1; i >= 0; i-- {
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
