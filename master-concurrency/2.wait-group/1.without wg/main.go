package main

import (
	"fmt"
	"time"
)

func main() {

	for i:=0;i<10;i++{
		go work(i+1)
	}
	time.Sleep(100*time.Millisecond)
	fmt.Println("main exiting")
}

func work(id int) {
	time.Sleep(100*time.Millisecond)
	fmt.Println("done with ",id," id")
}