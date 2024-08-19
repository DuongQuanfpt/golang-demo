package main

import (
	"fmt"

)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if(x<0){
		var e ErrNegativeSqrt = ErrNegativeSqrt(x)
		return x, e
	}	
	
	z := 1.0
	for i := 0 ; i <=10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z,nil
	
}



type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func interfaceTutorial() {
	shapes:= []Shape{Rectangle{6, 8}, Square{5}}

	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)//(<nil>, <nil>)
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)//(42, int)
	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)//(hello, string)
	
	for _, shape := range shapes {
		fmt.Printf("%v %T \n",shape ,shape)
		//{6 8} main.Rectangle  {5} main.Square
	}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
		//48 25
	}

	var shapeI Shape = Rectangle{6,8}
	t, ok := shapeI.(Rectangle)
	fmt.Println(t , ok)//{6 8} true
	shapeI = Square{5}
	t, ok = shapeI.(Rectangle)

	shapeI = Square{5}
	switch v:= shapeI.(type) {
	case Rectangle:
		fmt.Printf("Rectangle with length: %v width: %v",v.length,v.width)
	case Square:
		fmt.Printf("Square with side: %v",v.side)
	}

	
}


