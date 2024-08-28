package main

import (
	"fmt"
	"unicode/utf8"
)
func dataTypeTutorial(){
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum float32 = 12345678 / 9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum8 int8 = 2
	var result float32 = floatNum32 + float32(intNum8)
	fmt.Println(result)

	var intNum1 int = 5
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString(myString))

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)
	
	var1,var2 := 1,2
	fmt.Println(var1,var2)

	const myConst string = "const value"
	fmt.Println(myConst)
}