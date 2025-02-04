package sprint

func ToThePowerRecursive(n, power int) int {
    if power < 0 {
        return 0
    }

    if power == 0 {
        return 1
    }

    return n * ToThePowerRecursive(n, power-1)
}
