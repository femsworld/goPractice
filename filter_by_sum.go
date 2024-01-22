package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	var result [][]int

	for _, subarray := range arr {
		sum := calculateSum(subarray)
		if sum >= limit {
			result = append(result, subarray)
		}
	}

	return result
}

func calculateSum(subarray []int) int {
	sum := 0
	for _, value := range subarray {
		sum += value
	}
	return sum
}
