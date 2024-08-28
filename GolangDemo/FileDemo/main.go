package main

import (
	"log"
	"os"
)

func main() {
	//create file
	// if err := os.MkdirAll("a/b/c/d/hello.txt", os.ModePerm); err != nil {
	// 	log.Fatal(err)
	// }

	//write file
	err := os.WriteFile("hello.txt", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}