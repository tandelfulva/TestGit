package main

import "fmt"

func main() {
	//defer close()
	//panic
	//Open Resource
	defer func() {
		fmt.Println("Close Resources")
		details := recover()
		fmt.Println("Details:", details)
	}()

	fmt.Println("Open Resources: Hello,I am Fulva")
	panic("Panicking: I don't know what to do!")
	//recover()
	fmt.Println("End of main")
}
