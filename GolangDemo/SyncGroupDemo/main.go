package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	maxGoroutines := 15
	guard := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		guard <- struct{}{} // would block if guard channel is already filled
		wg.Add(1)
		go func(n int) {
			defer fmt.Println("idk")
			worker(n)
			defer func(){<-guard}()
			defer wg.Done()
			if(true){
				return
			}
			
		}(i)
	}
	wg.Wait()
}

func worker(i int) {
	fmt.Println("doing work on", i)
	time.Sleep(5 * time.Second)
	fmt.Println("Done at", i)
}