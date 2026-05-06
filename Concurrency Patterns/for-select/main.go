package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		i :=0
		for i < 100{
			fmt.Println("csk")
			i++
		}
	}()
	go func() {
		i := 0
		for i<100{
			fmt.Println("rcb")
			i++
		}
	}()

	time.Sleep(2 * time.Second)

}