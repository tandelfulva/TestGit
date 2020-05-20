package main

import "fmt"

func mainFunction(name string) func() {
	return func() {
		secondFunction(name)
	}
}

func secondFunction(name string) {
	fmt.Println(name)
}

func main() {
	value := mainFunction("Hello,I am Fulva")
	value()
}
