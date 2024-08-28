package main

import (
	"fmt"
	"strings"
)

func runesTutorial() {
	var myString = []rune("résumé")
	var index = myString[0]
	fmt.Printf("stuff : %v , %T \n", index, index)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("the length of myString : %v \n", len(myString))

	var strSlice = []string{"a", "b", "c", "d"}
	var strSlice2 = "nothing "
	// create a new string =>bad
	for i := range strSlice {
		strSlice2 += strSlice[i]
	}
	fmt.Printf(strSlice2)
	// no new string created => good
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var strSlice3 = strBuilder.String()
	fmt.Printf("\n%v\n", strSlice3)

	var var1 complex64 = complex(3,-5)  
	
	fmt.Println("var1:", var1)
	var z1 complex64 = 1 + 2i
    var z2 complex128 = 3 + 4i

    fmt.Println(z1, z2)
}
