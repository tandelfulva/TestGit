package main

import (
	"fmt"

	"github.com/tandelfulva/TestGit/myp"
	"github.com/tandelfulva/TestGit/mypack"
	"github.com/tandelfulva/TestGit/packages/addition"
)

func main() {
	myp.MyFun1()
	mypack.MyFunction()
	fmt.Println("Additon is: ", addition.Add(52, 12))
}
