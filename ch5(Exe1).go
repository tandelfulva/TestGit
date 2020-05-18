package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [5]string{"f", "u", "l", "v", "a"}

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
}
