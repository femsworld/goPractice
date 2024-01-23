package sprint

func CombN(n int) []string {
	var result []string
	generateCombination("", 0, n, &result)
	return result
}

func generateCombination(current string, start, n int, result *[]string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}

	for i := start; i <= 9; i++ {
		newCombination := current + string('0'+i)
		generateCombination(newCombination, i+1, n-1, result)
	}
}
