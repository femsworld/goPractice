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
	minimumToAdd := 0
	if trueCount > falseCount {
		minimumToAdd = trueCount - falseCount
	} else if falseCount > trueCount {
		minimumToAdd = falseCount - trueCount
	}

	// Add the minimum number of booleans necessary
	for i := 0; i < minimumToAdd; i++ {
		arr = append(arr, !arr[0])
	}

	return arr
}
