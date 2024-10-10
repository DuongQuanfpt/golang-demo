package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// maxGoroutines := 15
	// guard := make(chan struct{}, maxGoroutines)
	// var wg sync.WaitGroup

	// for i := 0; i < 30; i++ {
	// 	guard <- struct{}{} // would block if guard channel is already filled
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		defer fmt.Println("idk")
	// 		worker(n)
	// 		defer func(){<-guard}()
	// 		defer wg.Done()
	// 		if(true){
	// 			return
	// 		}

	// 	}(i)
	// }
	// wg.Wait()

	var wg sync.WaitGroup
	guard := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			guard <- i
		}
		close(guard)
	}()

	sum := []int{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range guard {
			sum = append(sum, v) // Directly append here without spawning additional goroutines
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range guard {
			sum = append(sum, v) // Directly append here without spawning additional goroutines
		}
	}()
	wg.Wait()
	fmt.Println(sum)
}

func worker(i int) {
	fmt.Println("doing work on", i)
	time.Sleep(5 * time.Second)
	fmt.Println("Done at", i)
}
