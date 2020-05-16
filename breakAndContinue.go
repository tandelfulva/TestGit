package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		fmt.Println("Iterating ", i)

		if i == 6 {
			continue
		}

		if i == 8 {
			break
		}

		if i == 4 {
			return
		}
	}
	fmt.Println("Outside For Loop")

}
