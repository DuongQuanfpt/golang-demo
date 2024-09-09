package main

import (
	// "fmt"
	"sync"
	
)

type Acc struct {
	mu sync.Mutex
	balance int
}

func (a *Acc) withdraw(amt int)  {
	a.mu.Lock()
	a.balance -= amt
	a.mu.Unlock()
}

func (a *Acc) deposit(amt int)  {
	a.mu.Lock()
	a.balance += amt
	a.mu.Unlock()
}

func sendIntToChannel(i []int, c chan int) {
}

func sum(a int, b int, c chan int) {
	c <- a + b
}

func concurrencyTutorial() {
	// result := make(chan int, 0)
	// go sum(5,5,result)
	// go sum(5,8,result)

	// value := <-result
	// fmt.Println(value)
	// value = <-result
	// fmt.Println(value)

	ch := make(chan int)
	go func ()  {
		ch <- 1
		ch <- 2
	}()
	
	// ch <- 2
	// close(ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	
	// c1 := make(chan int)
	// c2 := make(chan string)
	// exit := make(chan int)
	
	// go func() {
	// 	intS := []int{1,2,3,4,5}
	// 	for i:= range intS {
	// 		c1 <- intS[i]
	// 	}
	// 	exit <- 1
	// }()
	// go func() {
	// 	stringS := []string{"one","two","three","four","five"}
	// 	for  i:= range stringS{
	// 		c2 <- stringS[i]
	// 	}
	// 	exit <- 1
	// }()
	
	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println("received", msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println("received", msg2)
	// 	case <-exit:
	// 		fmt.Println("Done! exitted")
	// 		return
	// 	default:
	// 		fmt.Println("waiting...")
	// 	}
	// }
	
	// account := Acc{balance: 100}

	// for i := 0; i < 1000; i++ {
	// 	go account.deposit(10)
	// }

	// for i := 0; i < 1000; i++ {
	// 	go account.withdraw(10)
	// }

	// time.Sleep(2* time.Second)
	// fmt.Println(account.balance)//100
}

