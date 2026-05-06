package main

import (
	"fmt"
	"time"
)

func channel(){
	myChannel:=make(chan string)

	go func ()  {
		myChannel <- "data"
	}()
	result:= <-myChannel
	fmt.Println(result)
}


// channel()