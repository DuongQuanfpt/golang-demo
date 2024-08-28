package main

import (
	"fmt"
	"regexp"
)

func main() {
	// a regex object which
	// can be reused later
	re, _ := regexp.Compile("geek")

	// string to be matched
	str := "I love geeksforgeeks"

	// returns the slice of first
	// and last index
	match := re.FindAllIndex([]byte(str),-1)
	fmt.Println(match)
	fmt.Println(str[7:11])

}