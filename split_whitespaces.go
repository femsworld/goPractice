package sprint

func isSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhitespaces(s string) []string {
	var words []string
	currentWord := ""

	for i := 0; i < len(s); i++ {
		if isSpace(s[i]) {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(s[i])
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
