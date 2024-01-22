package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	length := len(arr)
	from = normalizeIndex(from, length)
	to = normalizeIndex(to, length)

	// Swap indices if from is greater than or equal to to
	if from >= to {
		from, to = to, from
	}

	// return append(arr[:from], arr[to+1:]...)
	return append(arr[:from], arr[to:]...)
}

func normalizeIndex(index, length int) int {
	if index < 0 {
		return length + index
	}
	if index >= length {
		return length - 1
	}
	return index
}
