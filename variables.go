package main

import (
	"fmt"
	"reflect"
	"runtime"
)


var num = 52
//declare but not used
//Not give error
var enrollNo int

func main(){
	//var no=52
	//Give Error Declare but not used
	enrollNo = 160330107052
	var name string
	name = "Fulva"
	name = "Sonu"
	fmt.Println(reflect.TypeOf(enrollNo))
	fmt.Println(name)
	fmt.Printf("%T",name)
	fmt.Println()
	fmt.Printf("%T",enrollNo)
	fmt.Println()
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	//fmt.Println(num)
}