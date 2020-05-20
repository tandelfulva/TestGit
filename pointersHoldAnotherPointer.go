package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	fmt.Println("Value of a: ", a)
	fmt.Println("Memory address of a: ", &a)
	var p *int = &a
	fmt.Println("Value of P: ", p)
	fmt.Println("Dereferencing of P: ", *p)
	fmt.Println("Type of P: ", reflect.TypeOf(p))
	fmt.Println("Memory address of p: ", &p)

	var pointerToPointer **int = &p
	fmt.Println("Value of pointerToPointer: ", pointerToPointer)
	fmt.Println("Dereferencing of pointerToPointer: ", *pointerToPointer)
	fmt.Println("Double Dereferencing of pointerToPointer: ", **pointerToPointer)
	fmt.Println("Type of pointerToPointer: ", reflect.TypeOf(pointerToPointer))
	fmt.Println("Memory address of pointerToPointer: ", &pointerToPointer)

}
