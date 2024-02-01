package sprint

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func removeWhitespaceAndToLower(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		if !isWhitespace(s[i]) {
			result = append(result, toLower(s[i]))
		}
	}
	return string(result)
}

func IsPalindrome(s string) bool {
	s = removeWhitespaceAndToLower(s)

	// Compare the string with its reverse
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
