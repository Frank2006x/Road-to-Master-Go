package main

import (
	"fmt"
	"time"
)

func someThing(n int){
	fmt.Print(n)
}




func main() {
	go someThing(2)
	go someThing(3)
	go someThing(4)
	time.Sleep(2*time.Second)
	
	fmt.Print("hello")	
}