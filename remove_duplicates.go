package sprint

func RemoveDuplicates(arr []int) []int {
	uniqueElements := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if !uniqueElements[num] {
			uniqueElements[num] = true
			result = append(result, num)
		}
	}

	return result
}
