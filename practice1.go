package main

import "fmt"

func main() {
	sum := 0
	for sum < 5 {
		sum += sum
	}
	fmt.Println(sum)
}
