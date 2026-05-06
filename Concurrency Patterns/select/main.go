package main

import (
	"fmt"
	"time"
)




func main(){
	myChannel:=make(chan string)
	anotherChannel:=make(chan string)
	go func ()  {
		time.Sleep(2 * time.Second)
		myChannel <- "data"
	}()
	go func () {
		anotherChannel <- "goat"
	}()
	var result string
	select {
		case result = <-myChannel:
			fmt.Println(result)
		case result = <-anotherChannel:
			fmt.Println(result)
	}


}