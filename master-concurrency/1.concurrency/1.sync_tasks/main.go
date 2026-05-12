package main

import (
	"fmt"
	"time"
)

func main() {
	start:=time.Now()
	task1()
	task2()
	task3()
	task4()
	

	fmt.Println("elapsed",time.Since(start))
}

func task1() {
	time.Sleep(100*time.Millisecond)
	fmt.Println("running task1")
}
func task2() {
	time.Sleep(200*time.Millisecond)
	fmt.Println("running tas2")
}
func task3() {
	fmt.Println("running task3")
}
func task4() {
	time.Sleep(100*time.Millisecond)
	fmt.Println("running task4")
}