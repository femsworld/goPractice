package sprint

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return 'z' - (ch - 'a')
	} else if ch >= 'A' && ch <= 'Z' { 
		lowercaseRune := ch + ('a' - 'A')
		return 'z' - (lowercaseRune - 'a')
	} else {
		return ch
	}
}