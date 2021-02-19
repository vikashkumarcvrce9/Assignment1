package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("enter your age:")

	var n string
	fmt.Scan(&n)
	rahul, _ := strconv.ParseInt(n, 0, 64)
	fmt.Println(rahul)

}
