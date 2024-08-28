package main

import (
	"fmt"
	"math"
)
 

 func adder() func(int) int {
	var sum = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func something(a *Employee)  {
	fmt.Println(a.address)
}

func fibonacci() func() int {
	num1 := 0
	num2 := 0
	return func() int {
		if num1 == 0 {
			num1 = 1
			return num1
		}

		num2 = num1 + num2
		return num2
	}
}

func functionValueClosure() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}


	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))//math.sqrt(3*3 + 4*4)
	fmt.Println(compute(math.Pow))//3^4
	
	positive , negative := adder(), adder()
	for i := 0; i < 10; i++ {	
		fmt.Printf("%v   %v \n",positive(1) , negative(-1) ) 
	}
}

