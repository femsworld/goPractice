package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	start, end := from, to
	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}

	copy(arr[start:], arr[end:])
	arr = arr[:len(arr)-(end-start)]

	return arr
}
