package sprint

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	decimalValue := ConvertToDecimal(nbr, baseFrom)
	return ConvertFromDecimal(decimalValue, baseTo)
}

func ConvertToDecimal(s, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := len(s) - 1; i >= 0; i-- {
		digit := theIndexOf(base, s[i])
		if digit == -1 {
			return 0
		}
		result += digit * thePower(baseLen, len(s)-1-i)
	}

	return result
}

func ConvertFromDecimal(n int, base string) string {
	if !isValidBase(base) {
		return "0"
	}

	if n == 0 {
		return string(base[0])
	}

	result := ""
	baseLen := len(base)

	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n /= baseLen
	}

	return result
}

func isValidBase(base string) bool {
	return len(base) >= 2 && !InvalidChars(base, "+-")
}

func InvalidChars(s string, invalidChars string) bool {
	for _, char := range s {
		if thatContains(invalidChars, char) {
			return true
		}
	}
	return false
}

func thatContains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func theIndexOf(s string, char byte) int {
	for i, c := range s {
		if byte(c) == char {
			return i
		}
	}
	return -1
}

func thePower(x, y int) int {
	result := 1
	for y > 0 {
		result *= x
		y--
	}
	return result
}
