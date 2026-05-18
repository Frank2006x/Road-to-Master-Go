package main

import "fmt"

func main() {

    input := make(chan int)
	output:=make(chan int)
    go func() {
        input <- 10
    }()

    go func() {

        res := <-input

        res = res*9 + 7

        output <- res
    }()

    fmt.Println(<-output)
}