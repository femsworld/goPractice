package sprint

func Swap(a []string, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func SortWordArr(a []string) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				Swap(a, i, j)
			}
		}
	}
	return a
}
