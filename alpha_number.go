package sprint

func AlphaNumber(n int) string {
	if n == 0 {
		return "a"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var numStr string
	for n > 0 {
		digit := n % 10
		numStr = string('a'+digit) + numStr
		n /= 10
	}
	if isNegative {
		numStr = "-" + numStr
	}
	return numStr
}
