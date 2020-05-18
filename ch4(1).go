package main

import (
	"calc"
	"fmt"
	"greeting"
)

func main() {
	greeting.Hello()
	greeting.Hii()

	fmt.Println(calc.Add(12, 52))
	fmt.Println(calc.Subtract(7, 3))
}
