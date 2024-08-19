package hello

import "fmt"

type Something struct{
	a int
	b string
} 

func SayHello(name string) {
	fmt.Println("Hello,", name)
}

func saySomething() {
	fmt.Println("something")
}

