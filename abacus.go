package main

import "fmt"

func Abacus(a, b int) int {
	result := a/b
	return result
}

func main() {
	result :=Abacus(8,3)
	fmt.Println(result)

}