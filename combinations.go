package sprint

import "fmt"

func Combinations() string {
	result := ""

	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				result += fmt.Sprintf("%d%d%d", i, j, k)
				if i != 7 || j != 8 || k != 9 {
					result += ", "
				}
			}
		}
	}
	return result
}
