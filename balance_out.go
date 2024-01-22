package sprint

func BalanceOut(arr []bool) []bool {
	trueCount, falseCount := countBooleans(arr)

	for trueCount != falseCount {
		if trueCount < falseCount {
			arr = append(arr, true)
			trueCount++
		} else {
			arr = append(arr, false)
			falseCount++
		}
	}

	return arr
}

func countBooleans(arr []bool) (int, int) {
	trueCount, falseCount := 0, 0

	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}
	return trueCount, falseCount
}
