package main

import "fmt"

func main() {
	addFun1 := sumOfNumbers(12, 52)
	addFun2, subFun2 := sumAndSubOfNumbers(96, 52)
	fmt.Println("Addition in function1:", addFun1)
	fmt.Println("Addition in function2:", addFun2)
	fmt.Println("Substraction in function2:", subFun2)

}

func sumOfNumbers(num1 int32, num2 int32) float64 {
	return float64(num1 + num2)
}

func sumAndSubOfNumbers(num1 int32, num2 int32) (float64, float64) {
	return float64(num1 + num2), float64(num1 - num2)
}
