package sprint

func ToCapitalCase(s string) string {
    runes := []rune(s)
    var result []rune
    capitalizeNext := true

    for _, char := range runes {
        if isAlphanumeric(char) {
            if capitalizeNext {
                result = append(result, toUpperCase(char))
                capitalizeNext = false
            } else {
                result = append(result, toLowerCase(char))
            }
        } else {
            result = append(result, char)
            capitalizeNext = true
        }
    }

    return string(result)
}

func isAlphanumeric(char rune) bool {
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toUpperCase(char rune) rune {
    if char >= 'a' && char <= 'z' {
        return char - 32
    }
    return char
}

func toLowerCase(char rune) rune {
    if char >= 'A' && char <= 'Z' {
        return char + 32
    }
    return char
}