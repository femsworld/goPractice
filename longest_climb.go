package sprint

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{} // Empty array case
	}

	startIndex := 0
	currentStart := 0
	maxLength := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if i-currentStart+1 > maxLength {
				startIndex = currentStart
				maxLength = i - currentStart + 1
			}
		} else {
			currentStart = i
		}
	}

	return arr[startIndex : startIndex+maxLength]
}
