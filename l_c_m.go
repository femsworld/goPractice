package sprint

func GCD1(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func LCM(a, b int) int {
    gcd := GCD1(a, b)
    lcm := (a * b) / gcd
    return lcm
}