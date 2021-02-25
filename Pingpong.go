package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; true; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println(i, " ", "ping")
		time.Sleep(time.Second * 2)
		fmt.Println(i, " ", "pong")

	}
}
