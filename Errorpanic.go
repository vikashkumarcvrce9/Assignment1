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
	var t float64
	fmt.Scanln(&t)

	result, err := sqrt(t)
	if err != nil {
		panic(err)       //here panic is Created
		fmt.Println(err) //If You Comment panic this Statemnt Next Line Also Printed
		fmt.Print("Statement belongs After panic")

	} else {
		fmt.Println(result)
	}

}
