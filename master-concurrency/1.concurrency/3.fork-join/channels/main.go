package main

import (
	"fmt"
	"time"
)

func main() {
	now:=time.Now()
	done:=make(chan struct{})	
	go func(){
		work()
		done<-struct{}{}
	}()
	<-done
	fmt.Println("done waiting exiting")
	fmt.Println(time.Since(now))
}

func work() {
	time.Sleep(300*time.Millisecond)

}