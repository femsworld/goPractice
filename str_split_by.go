package sprint

func StrSplitBy(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	var substrings []string
	startIndex := 0

	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			substrings = append(substrings, s[startIndex:i])
			startIndex = i + len(sep)
			i += len(sep) - 1
		}
	}

	if startIndex <= len(s) {
		substrings = append(substrings, s[startIndex:])
	}

	return substrings
}
