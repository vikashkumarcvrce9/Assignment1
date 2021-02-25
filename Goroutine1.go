package main

import (
	"fmt"
	"time"
)

func main() {
	go count("mohan")
	count("Rahul")
}
func count(name string) {
	for i := 1; i < 50; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Millisecond * 500)
	}
}
