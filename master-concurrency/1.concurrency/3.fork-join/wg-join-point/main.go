package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now:=time.Now()
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("done waiting exiting")
	fmt.Println(time.Since(now))
}

func work() {
	time.Sleep(300*time.Millisecond)

}