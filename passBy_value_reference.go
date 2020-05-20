package main

import "fmt"

func main() {
	number := 20
	case1(number)
	fmt.Println("After case1, value in main: ", number)
	fmt.Println("After case1, Address of number in main: ", &number)
	case2(&number)
	fmt.Println("After case2, value in main: ", number)
	fmt.Println("After case2, Address of number in main: ", &number)
}

//pass by value
func case1(number int) {
	number = number + 2
	fmt.Println("In case1, value is: ", number)
	fmt.Println("In case1, address is: ", &number)

}

//pass by reference

func case2(number *int) {
	*number = *number + 2
	fmt.Println("In case2, value is: ", number)
	fmt.Println("In case2, address is: ", &number)

}
