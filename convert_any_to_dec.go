package sprint

func ConvertAnyToDec(s string, base string) int {
	if len(base) < 2 || containsInvalidChars1(base, "+-") {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := len(s) - 1; i >= 0; i-- {
		digit := indexOf(base, s[i])
		if digit == -1 {
			return 0
		}
		result += digit * power(baseLen, len(s)-1-i)
	}

	return result
}

func containsInvalidChars1(base string, invalidChars string) bool {
	for _, char := range base {
		if contains(invalidChars, char) {
			return true
		}
	}
	return false
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func indexOf(s string, char byte) int {
	for i, c := range s {
		if byte(c) == char {
			return i
		}
	}
	return -1
}

func power(x, y int) int {
	result := 1
	for y > 0 {
		result *= x
		y--
	}
	return result
}
