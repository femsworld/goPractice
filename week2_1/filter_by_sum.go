package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	var result [][]int

	for _, subarray := range arr {
		sum := 0
		for _, value := range subarray {
			sum += value
		}

		if sum >= limit {
			result = append(result, subarray)
		}
	}

	if len(result) == 0 {
		return [][]int{}
	}

	return result
}
