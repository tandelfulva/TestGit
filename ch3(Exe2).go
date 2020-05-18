package main

import (
	"errors"
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't devide by zero")
	}
	return dividend / divisor, nil
}
func main() {
	quotient, err := divide(5.6, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
