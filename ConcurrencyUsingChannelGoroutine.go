package main

import (
	"fmt"
)

func add(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	c <- sum //send to channnel
}
func main() {
	s1 := []int{22, 1, 2, 8, 8, 111, 5, 7}
	s := []int{0, 1, 2, 3, 4, 5, 6}

	c := make(chan int)

	go add(s, c)
	go add(s1, c)

	x := <-c //recive to channel
	y := <-c

	fmt.Println(x)
	fmt.Println(y)

}
