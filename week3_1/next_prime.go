package sprint

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for i := n; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
