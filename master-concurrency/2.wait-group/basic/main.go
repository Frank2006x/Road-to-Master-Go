package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	
	go func(){
		defer wg.Done()
		time.Sleep(10*time.Millisecond)
		fmt.Println("1")
	}()
	go func(){
		defer wg.Done()
		time.Sleep(10*time.Millisecond)
		fmt.Println("2")
	}()
	go func(){
		defer wg.Done()
		time.Sleep(10*time.Millisecond)
		fmt.Println("3")
		
	}()

	wg.Wait()



}
