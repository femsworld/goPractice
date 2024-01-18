package sprint

func Doop(a int, op string, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			return 0
		}
	case "%":
		if b != 0 {
			return a % b
		} else {
			return 0
		}
	default:
		return 0
	}
}
