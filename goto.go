package main

import "fmt"

func main() {

	var num int
test:
	//lable
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The number is: ", num)

	if num > 200 {
		fmt.Println("Exiting")
		return

	}

	if num > 100 {
		fmt.Println("Number very high")
		goto test
	}
	fmt.Println("Good number")

}
