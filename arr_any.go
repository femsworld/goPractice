package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) {
			return true
		}
	}
	return false
}

func IsUpper(s string) bool {
	for _, char := range s {
		if !(char >= 'A' && char <= 'Z') {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if !(isLetter(char) || isDigit(char)) {
			return false
		}
	}
	return true
}
func isLetter(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
