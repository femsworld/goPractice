package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	from = (from + len(arr)) % len(arr)
	to = (to + len(arr)) % len(arr)

	if from > to {
		from, to = to, from
	}

	return append(arr[:from], arr[to:]...)
}
