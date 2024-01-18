package sprint

import "fmt"

func Pairs() string {
	result := ""

	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			result += fmt.Sprintf("%02d %02d", i, j)
			if i != 98 || j != 99 {
				result += ", "
			}
		}
	}
	return result
}
