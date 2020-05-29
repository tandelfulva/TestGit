package main

import "fmt"

func main() {
	myMap := map[int]string{
		52: "Fulva",
		32: "Pinkal",
		29: "Nensi",
	}
	fmt.Println("Before Changes my map is:", myMap)
	newMap := checkMap(myMap)
	fmt.Println("My Map is:", myMap)
	fmt.Println("New Map is:", newMap)
}

func checkMap(myMap map[int]string) map[int]string {
	myMap[52] = "Sonu"
	return myMap
}
