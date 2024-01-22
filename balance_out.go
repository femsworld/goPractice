package sprint

func BalanceOut(arr []bool) []bool {
	// Count the number of true and false values in the array
	trueCount, falseCount := 0, 0
	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}
	// Calculate the number of booleans needed to balance the array
	minimumToAdd := abs(trueCount - falseCount)

	// Add the minimum number of booleans necessary
	for i := 0; i < minimumToAdd; i++ {
		arr = append(arr, !arr[0]) // add the opposite of the first element
	}

	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
