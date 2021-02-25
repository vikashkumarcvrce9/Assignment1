package main

import (
	"fmt"
	"math/rand"
)

func num(a int, c chan int) {
	numb := rand.Intn(a)

	c <- numb
}
func main() {
	c := make(chan int)
	go num(100, c)
	x := <-c

	fmt.Println(x)
}
