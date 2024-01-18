package sprint

import "fmt"

func Pairs() string {
	result := ""

	for i := 10; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			// result += fmt.Sprintf("%02d %02d, ", i, j)
			result += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}
	// Remove the trailing ", " before returning the result
	return result[:len(result)-2]

}
