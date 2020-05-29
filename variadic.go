package main

import "fmt"

func main() {
	fmt.Println("Function with 0 argument: ", sum(""))
	fmt.Println("Function with 1 argument: ", sum("", 10))
	fmt.Println("Function with 2 argument: ", sum("", 10, 20))
	fmt.Println("Function with 3 argument: ", sum("", 10, 20, 30))
	var mySlice = []int{10, 20, 30, 40, 50}
	fmt.Println(sum("", mySlice...))

}

func sum(str string, num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total

}
