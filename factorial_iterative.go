package sprint

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}

	result := 1

	for i := 1; i <= n; i++ {

		if result > (1<<31-1)/i {
			return 0
		}

		result *= i
	}

	return result
}
