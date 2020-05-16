package main

import "fmt"

func main() {
	usernum := addOne(52)
	fmt.Println(usernum)
}

func addOne(num int) int {

	num = num + 1 // num=52+1=53
	return num

}
