package main

import "fmt"

func main() {
	var a int = 200
	var b int = 300
	fmt.Printf("Max value is: %d\n", max(a, b))
}
func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
