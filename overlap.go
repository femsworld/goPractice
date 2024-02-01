package sprint

func Overlap(arr1, arr2 []int) []int {
	result := []int{}
	frequencyMap := make(map[int]int)

	for _, num := range arr1 {
		frequencyMap[num]++
	}
	for _, num := range arr2 {
		if count, ok := frequencyMap[num]; ok && count > 0 {
			result = append(result, num)
			frequencyMap[num]--
		}
	}
	bubbleSort(result)

	return result
}
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
