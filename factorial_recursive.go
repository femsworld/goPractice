package sprint

func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	subFactorial := FactorialRecursive(n - 1)

	result, overflow := safeMultiply(n, subFactorial)
	if overflow {
		return 0
	}

	return result
}

func safeMultiply(a, b int) (int, bool) {
	result := a * b

	if b != 0 && result/b != a {
		return 0, true
	}

	return result, false
}
