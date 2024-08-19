package main

import "fmt"

type Person struct {
	name    string
	age     string
	phone   string
	address string
}

func structTutorial() {
	var person Person = Person{"Quan", "24", "phone", "address"}
	personP := &person

	fmt.Println(personP.address)
}
