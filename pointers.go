package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b string = "Fulva"
	fmt.Println("Value of a: ", a)
	fmt.Println("Memory address of a: ", &a)
	p := &a
	q := &b
	fmt.Println("Value of P: ", p)
	fmt.Println("Dereferencing of P: ", *p)
	fmt.Println("Type of P: ", reflect.TypeOf(p))
	fmt.Println("Memory address of p: ", &p)

	fmt.Println("Value of q: ", q)
	fmt.Println("Dereferencing of q: ", *q)
	fmt.Println("Type of q: ", reflect.TypeOf(q))
	fmt.Println("Memory address of q: ", &q)
}
