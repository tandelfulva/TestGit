package main

import "fmt"

func main() {
	displayName()
	fmt.Println("End of main")
	displayName()
}

func displayName() {
	fmt.Println("Fulva")
}
