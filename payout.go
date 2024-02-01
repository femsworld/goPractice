package sprint

import "sort"

func Payout(amount int, denominations []int) []int {
	if amount < 0 || len(denominations) == 0 || !isPositiveDenominations(denominations) {
		return []int{} // Invalid input cases
	}

	sort.Sort(sort.Reverse(sort.IntSlice(denominations))) // Sort denominations in descending order

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

	return []int{} // Return an empty array if the payout cannot be made
}

func isPositiveDenominations(denominations []int) bool {
	for _, denom := range denominations {
		if denom <= 0 {
			return false
		}
	}
	return true
}
