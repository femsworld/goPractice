package sprint

func PascalsTriangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}

	return triangle
}
