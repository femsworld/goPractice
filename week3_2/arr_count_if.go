package sprint

func IsLower(s string) bool {
    for _, r := range s {
        if r < 'a' || r > 'z' {
            return false
        }
    }
    return true
}

func IsUpper(s string) bool {
    for _, r := range s {
        if r < 'A' || r > 'Z' {
            return false
        }
    }
    return true
}

func IsNumeric(s string) bool {
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}

func IsAlphanumeric(s string) bool {
    for _, r := range s {
        if !IsLower(string(r)) && !IsUpper(string(r)) && !IsNumeric(string(r)) {
            return false
        }
    }
    return true
}

func ArrCountIf(f func(string) bool, tab []string) int {
    count := 0
    for _, s := range tab {
        if f(s) {
            count++
        }
    }
    return count
}