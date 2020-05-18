package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 52
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
