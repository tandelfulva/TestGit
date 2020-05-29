package main

import "fmt"

func main() {
	var person1 = &person{
		firstName: "Fulva",
		lastName:  "Tandel",
		age:       22,
	}
	fmt.Println("Details of Person1 :", *person1)
	fmt.Println("Address is:", &person1)
	fmt.Println("First Name of Person1:", person1.firstName)
	fmt.Println("Last NAme of Person1:", person1.lastName)
	fmt.Println("Age of Person1:", person1.age)
	fmt.Println()

	var person2 = newPerson("Sonu", "Tandel", 21)
	fmt.Println("Person2 is:", *person2)
	fmt.Println("Address is:", &person2)
}

func newPerson(fName, lName string, age int) *person {
	var personTest = &person{
		firstName: fName,
		lastName:  lName,
		age:       age,
	}
	return personTest

}

type person struct {
	firstName string
	lastName  string
	age       int
}
