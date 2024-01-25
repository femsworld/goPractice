package sprint


func NbrBase(n int, base string) string {
    if len(base) < 2 || containsInvalidChars(base) || hasRepeatingChars(base) {
        return "NV"
    }

    isNegative := n < 0
    n = abs(n)

    result := ""
    baseLen := len(base)

    for n > 0 {
        remainder := n % baseLen
        result = string(base[remainder]) + result
        n /= baseLen
    }

    if isNegative {
        result = "-" + result
    }

    return result
}

func containsInvalidChars(base string) bool {
    for _, char := range base {
        if char == '+' || char == '-' {
            return true
        }
    }
    return false
}

func hasRepeatingChars(base string) bool {
    charSet := make(map[rune]bool)
    for _, char := range base {
        if charSet[char] {
            return true
        }
        charSet[char] = true
    }
    return false
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}