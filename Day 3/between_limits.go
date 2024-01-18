package sprint

func BetweenLimits(from, to rune) string {
	var result string
	start := from
	end := to
	if from > to {
		start, end = to, from
	}
	for r := start + 1; r < end; r++ {
		result += string(r)
	}
	return result
}
