package sprint

func Sqrt(n int) int {
	result := intSqrt1(n)
	if result*result == n {
		return result
	}
	return 0
}

func intSqrt1(n int) int {
	if n < 0 {
		return 0
	}
	result := 0
	for i := 0; i*i <= n; i++ {
		result = i
	}
	return result
}
