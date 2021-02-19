package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Negative nuumber passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	fmt.Println("Enter a Number")
	var a float64
	fmt.Scanln(&a)
	result, err := sqrt(a)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}

}
