package sprint

func AlphabetMastery(n int) string {
	if n < 1 || n > 26 {
		return ""
	}
	// letter := string('A' - 1 + n)
	output := ""
	for i := 0; i < n; i++ {
		output += string('a' + i)
	}
	return output
	// return strings.ToLower(letter)
}
