package main

import "fmt"

func main() {
	fmt.Println(saveDevide(10, 11))
	panic("Programme got panic ")
	fmt.Println(saveDevide(10, 2))
}
func saveDevide(a, b int) int {

	return a + b

}
