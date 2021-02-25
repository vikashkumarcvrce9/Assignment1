package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		count("sheep")
		wg.Done()
	}()
	wg.Wait()
	go func() {
		count("Rahul")
		wg.Done()
	}()
	wg.Wait()

}
func count(name string) {
	for i := 1; i < 6; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Millisecond * 500)
	}
}
