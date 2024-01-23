package sprint

func NRune(s string, i int) rune {
	if i >= 0 && i < len(s) {
		return []rune(s)[i]
	}
	return 0
}
