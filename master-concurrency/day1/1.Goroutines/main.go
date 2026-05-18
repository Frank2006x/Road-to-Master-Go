package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("hello from goroutine")
}

func main() {

    go hello()

    time.Sleep(time.Second)

    fmt.Println("main finished")
}