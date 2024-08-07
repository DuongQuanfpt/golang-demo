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
	fmt.Printf("\n%v", strSlice3)
}
