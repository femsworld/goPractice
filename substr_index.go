package sprint

func SubstrIndex(s string, toFind string) int {
	sLen, toFindLen := len(s), len(toFind)

	for i := 0; i <= sLen-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}
	return -1
}
