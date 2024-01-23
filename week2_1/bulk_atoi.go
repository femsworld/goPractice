package sprint

func SimpleStrToInt(s string) int {
	result := 0
	sign := 1

	for _, char := range s {
		switch {
		case char == '-':
			sign = -1
		case char >= '0' && char <= '9':
			digit := int(char - '0')
			result = result*10 + digit
		}
	}

	return result * sign
}

func BulkAtoi(arr []string) []int {
	result := make([]int, len(arr))

	for i, str := range arr {
		result[i] = SimpleStrToInt(str)
	}

	return result
}
