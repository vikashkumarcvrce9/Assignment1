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
	result, err := sqrt(-64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}
	result, err = sqrt(64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}

}
