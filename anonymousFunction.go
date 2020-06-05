package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is code inside anonymous function:")
		fmt.Println("Hello,I am Fulva")
		fmt.Println()
	}()
	anonyFunc()
}

func anonyFunc() {
	var num = func(s string) int {
		fmt.Println(s, "This is code inside non-main anonymous function")
		return 52
	}("Hello Fulva,")
	fmt.Println()
	fmt.Println("The Return value is:", num)
}
