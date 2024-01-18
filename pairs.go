package sprint

import "fmt"

func Pairs() string {
	result := ""

	for i := 10; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			// result += fmt.Sprintf("%02d %02d, ", i, j)
			result += fmt.Sprintf("%02d %02d", i, j)
			if i != 98 || j != 99 {
				result += ", "
			}
		}
	}
	// Remove the trailing ", " before returning the result
	return result
}
