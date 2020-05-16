package main

import "fmt"

func main() {
	cgpa := 9.32
	fmt.Printf("%T", cgpa)

	switch cgpa {
	case 10:
		fmt.Println("Excellent")

	case 9, 9.32:
		fmt.Println("Very Good")

	case 10000:
		fmt.Println("test")

	default:
		fmt.Println("Inside Default")

	}
}
