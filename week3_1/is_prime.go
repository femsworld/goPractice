package sprint

func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    sqrtN := intSqrt(n)
    for i := 2; i <= sqrtN; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func intSqrt(n int) int {
    x := n
    y := 1

    for x > y {
        x = (x + y) / 2
        y = n / x
    }
    return x
}

func printResult(result bool) {
    if result {
    } else {
    }
}

