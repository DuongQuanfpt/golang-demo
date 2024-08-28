package hello

import (
	"fmt"
)

var c, python, java bool

func func1() {
	var i int
	fmt.Println(i, c, python, java)
}

func func2() {
	fmt.Println(c, python, java)
}

func CompareSum(x, n, lim int)  {
	if sum := x+n; sum < lim {
		fmt.Printf("%v < %v\n", sum, lim)
	} else {
		fmt.Printf("%v >= %v\n", sum, lim)
	}
}
