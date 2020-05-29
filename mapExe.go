package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	for i := 1; i <= 5; i++ {
		myMap[i] = i * 'a'
	}
	fmt.Println("Befor changes myMap is:", myMap)
	fmt.Println()

	newMap := checkMap(myMap)
	fmt.Println("After changes myMap is:", myMap)
	fmt.Println("After changes newMap is:", newMap)
	fmt.Println()

	delete(myMap, 3)
	fmt.Println("After deleting myMap[3] myMap is:", myMap)
	fmt.Println("After deleting myMap[3] newMap is:", newMap)
	fmt.Println()

	var val, ok = myMap[3]
	fmt.Println("Value of myMap[3] is:", val)
	fmt.Println("Ok value is:", ok)
	fmt.Println()

	var value, okk = myMap[4]
	fmt.Println("Value myMap[4] is:", value)
	fmt.Println("Ok value is:", okk)
}

func checkMap(myMape map[int]int) map[int]int {

	for i := 1; i <= 5; i++ {
		myMape[i] = i * 'A'
	}
	return myMape
}
