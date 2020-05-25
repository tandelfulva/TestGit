package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray = [4]string{"Fulva", "Harjivandas", "Tandel"}
	fmt.Println(reflect.TypeOf(myArray))
	case1(myArray)
	fmt.Println("After case1 In main: ", myArray)
	for index, value := range myArray {
		fmt.Println("Index is: ", index, " Value is: ", value)
	}

	case2(&myArray)
	fmt.Println("After case2 In main: ", myArray)
	for index, value := range myArray {
		fmt.Println("Index is: ", index, " Value is: ", value)
	}

}

func case1(arr [4]string) {
	arr[0] = "sonu"
	fmt.Println("In case1: ", arr)

}

func case2(arr *[4]string) {
	arr[0] = "sonu"
	fmt.Println("In case2 (deref) : ", *arr)
	fmt.Println("In case2 : ", arr)
}
