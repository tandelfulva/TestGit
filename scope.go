package main

import "fmt"

var num int

func main() {
	display()
	num = 40
	//display()
	fmt.Println("in main: ", num)
	for i := 10; i < 14; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)   give error undefined variable i
}

func display() {
	var num1 int
	num1 = 30
	fmt.Println(num1)
	fmt.Println("in display: ", num)
}
