package main

import "fmt"

func main() {
	var num = 52
	function1()
	defer function2(num)
	num = 160330107052
	fmt.Println("End of main. Number is:", num)
}

func function1() {
	defer fmt.Println("Hello from Function1")
	fmt.Println("End of Function1")
}

func function2(num int) {
	fmt.Println("End of Function2. Number is:", num)
}
