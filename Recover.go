package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number 1:")
	var n1 int
	var n2 int
	fmt.Scanln(&n1)
	fmt.Println("Enter the Second number")
	fmt.Scanln(&n2)
	fmt.Println(SaveDivide(n1, n2))
}
func SaveDivide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
