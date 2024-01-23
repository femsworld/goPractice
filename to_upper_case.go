package sprint

func ToUpperCase(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char - 'a' + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
