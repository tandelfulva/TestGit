package main

import "fmt"

func add(val1 int32, val2 int32) int32 {
	return val1 + val2
}

func subtract(val1 int32, val2 int32) int32 {
	return val1 - val2
}

func mathOperations(testFun func(int32, int32) int32) int32 {
	var a, b int32
	a = 52
	b = 12
	return testFun(a, b)
}

func main() {
	addition := mathOperations(add)
	subtraction := mathOperations(subtract)
	fmt.Println("Addition is: ", addition)
	fmt.Println("subtraction is: ", subtraction)
}
