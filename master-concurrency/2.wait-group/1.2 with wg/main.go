package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup;
	start:=time.Now()
	wg.Add(10)
	for i:=range(10){
		go work(i+1,&wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))

	time.Sleep(100*time.Millisecond)
	fmt.Println("main exiting")
}

func work(id int,wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100*time.Millisecond)
	fmt.Println("done with ",id," id")
}