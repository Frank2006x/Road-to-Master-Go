package main

import (
	"fmt"
	"time"
)

func main() {
	start:=time.Now()
	done:=make(chan struct{})
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	<-done
	<-done
	<-done
	<-done

	fmt.Println("elapsed",time.Since(start))
}

func task1(done chan struct{}) {
	time.Sleep(100*time.Millisecond)
	fmt.Println("running task1")
	done <- struct{}{}
}
func task2(done chan struct{}) {
	time.Sleep(200*time.Millisecond)
	fmt.Println("running task2")
	done <- struct{}{}
}
func task3(done chan struct{}) {
	fmt.Println("running task3")
	done <- struct{}{}
}
func task4(done chan struct{}) {
	time.Sleep(100*time.Millisecond)
	fmt.Println("running task4")
	done <- struct{}{}
}