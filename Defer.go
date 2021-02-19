package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hi vikash ")
	for i := 0; i < 22; i++ {

		fmt.Println(i)
		defer fmt.Println(i)
		if i == 5 {
			panic("panic occurs")

		}

	}
}
