package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan int,5)
	sq:=make(chan int,5)
	start:=time.Now()
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	go func(){
		for v:=range input{
			sq <- v*v
		}
		close(sq)
	}()
	
		
	for v := range sq {
		fmt.Println(v)
	}

	fmt.Println(time.Since(start))
	

}