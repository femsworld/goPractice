package sprint

func Mean(a, b, c float32) float32 {
	result := (a + b + c)/3
	return result
}





// func Mean(numbers ...float32) float32 {
//     sum := float32(0)
//     count := len(numbers)

//     for _, num := range numbers {
//         sum += num
//     }

//     if count > 0 {
//         mean := sum / float32(count)
//         return mean
//     }

//     return 0
// }

