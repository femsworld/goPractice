package sprint

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return 'z' - (ch - 'a')
	} else {
		return ch
	}
}