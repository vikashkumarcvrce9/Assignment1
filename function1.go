package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var value int
	value = max(a, b)
	fmt.Println("Maximum vale in both is :-", value)
}
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
