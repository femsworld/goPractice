package sprint

func Countdown(n int) string {
	if n < 0 {
		return "0!"
	}
	result := ""
	for i := n; i >= 0; i -= 2 {
		if i < n {
			result += ", "
		}
		numStr := string('0' + byte(i%10))
		result += numStr
	}
	if n%2 == 1 {
		result += ", 0!"
	} else {
		result += "!"
	}
	return result
}
