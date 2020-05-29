package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	myMap["Fulva"] = 80
	myMap["Sonu"] = 90
	myMap["xyz"] = 100
	fmt.Println(myMap)
	fmt.Println()

	delete(myMap, "xyz")

	var score, ok = myMap["Fulva"]
	fmt.Println("Ok value is: ", ok)
	fmt.Println("Score of Fulva is: ", score)

	fmt.Println()

	var sc, okk = myMap["xyz"]
	fmt.Println("Okk value is: ", okk)
	fmt.Println("Score of xyz is: ", sc)

	fmt.Println()

	for key, val := range myMap {
		fmt.Println("Key is:", key, " Value is:", val)
	}

}
