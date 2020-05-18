package main

import (
	"fmt"
	"strings"
)

func main() {
	str, no := doesSomething("Fulva", "Tandel")
	fmt.Println(str, no)
}

func doesSomething(firstName string, lastName string) (string, int) {

	firstUpper := strings.ToUpper(firstName)
	SecondUpper := strings.ToUpper(lastName)
	fullName := firstUpper + " " + SecondUpper
	size := len(fullName)
	return fullName, size

}
