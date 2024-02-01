package sprint

func StrCompare(a, b string) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}

func IsSorted(f func(a, b string) int, arr []string) bool {
    n := len(arr)
    ascending := false
    descending := false

    for i := 0; i < n-1; i++ {
        cmp := f(arr[i], arr[i+1])

        if cmp > 0 {
            descending = true
        } else if cmp < 0 {
            ascending = true
        }

        if ascending && descending {
            return false
        }
    }
    return ascending || descending
}