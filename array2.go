package main

import "fmt"

func main() {
	var myArray = [...] int {12,03,1999}
	fmt.Println(myArray)
	var copyMyArray = myArray
	copyMyArray[0] = 40
	fmt.Println("First Element in myArray is: ",myArray[0])
	fmt.Println("First Element in copyMyArray is: ",copyMyArray[0])
	fmt.Println("My Array is: ",myArray)
	fmt.Println("Copy Array is: ",copyMyArray)
}