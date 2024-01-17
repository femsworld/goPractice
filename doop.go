package sprint

import "fmt"

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
			fmt.Println("Error: Division by zero")
			return 0
		}
	case "%":
		if b != 0 {
			return a % b
		} else {
			fmt.Println("Error: Modulo by zero")
			return 0
		}
	default:
		fmt.Println("Error: Invalid operator")
		return 0
	}
}
