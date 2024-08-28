package main

import (
	"fmt"
)

// import "fmt"

func pointerTutorial() {
	var intA int = 2
	var pint *int = &intA
	updateIntWithoutP(*pint)
	fmt.Println(intA)
	updateInt(pint)
	fmt.Println(intA)
}

func updateInt(a *int) {
	*a++
}

func updateIntWithoutP(a int)  {
	a++
}
