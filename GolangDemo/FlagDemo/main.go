package main

import (
	"flag"
	"fmt"
)

func main() {
	var taskNum int
	flag.IntVar(&taskNum, "n", 2, "pick (1-4) which task to run")
	flag.Parse()

	fmt.Println(taskNum,Sth)
}