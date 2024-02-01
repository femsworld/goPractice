package sprint

func LongestCommonSubstr(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)

	if len1 == 0 || len2 == 0 {
		return ""
	}
	lcs := make([][]int, len1+1)
	for i := range lcs {
		lcs[i] = make([]int, len2+1)
	}
	maxLength := 0
	endIndex := 0

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1

				if lcs[i][j] > maxLength {
					maxLength = lcs[i][j]
					endIndex = i - 1
				}
			}
		}
	}

	if maxLength > 0 {
		return str1[endIndex-maxLength+1 : endIndex+1]
	}

	return ""
}
