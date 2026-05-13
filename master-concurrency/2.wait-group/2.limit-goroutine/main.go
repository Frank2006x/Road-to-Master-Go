package main

import (
	"fmt"
	"sync"
)

type request func()

func main() {
	requests := map[int]request{}

	for i:=0;i<100;i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("Request",n)
			}
		}
		requests[i]=f(i)
	}

	var wg sync.WaitGroup

	max:=10

	for i:=0;i<100;i+=max{
		wg.Add(10)
		for j:=i;j<i+max;j++{
			go func (r request)  {
				defer wg.Done()
				r()
			}(requests[j])
		}
		fmt.Println("Patch")

		wg.Wait()
	}

}