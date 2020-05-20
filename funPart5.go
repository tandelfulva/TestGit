package main

import "fmt"

func mainFunction() func() {
	return func() {
		secondFunction()
	}
}

func secondFunction() {
	fmt.Println("Hello,I am Fulva")
}

func main() {
	value := mainFunction()
	value()
}
