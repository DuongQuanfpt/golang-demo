package main

import (
	"fmt"
	// "golang-tutorial/hello"
	"unsafe"
)

type Employee struct {
	name    string
	address string
	age     int
}

// func (a hello.Something) somethingFUnc(){

// }

func (a *Employee) updateName(newName string) {
	a.name = newName
}

func (a *Employee) showWithPointer() {
	fmt.Println(unsafe.Sizeof(a))
}

func (a Employee) show() {
	fmt.Println(unsafe.Sizeof(a))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return -float64(f)
	}
	return float64(f)
}

func methodTurorial() {
	person := Employee{
		name:    "Quan",
		address: "Somewhere",
		age:     24,
	}
	person.updateName("Duong Quan")
	fmt.Println(person.name) //Duong Quan

	person.show()//8
	person.showWithPointer()//40
}
