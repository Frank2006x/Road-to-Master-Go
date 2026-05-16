package main

import (
	"fmt"
	"sync"
)

func generator(start, end int) <-chan int {
	out := make(chan int)
	go func(){

		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func cube(c <-chan int) <-chan int {
	out := make(chan int)
	go func ()  {
		defer close(out)
		for v := range c {
			out <- v * v * v
		}
	}()
	return out
}

func merger(chal ...<-chan int)<-chan int {
	var wg sync.WaitGroup
	out:=make(chan int)
	wg.Add(len(chal))
	output := func (id <-chan int)  {
		defer wg.Done()
		for v:=range id{
			out<-v;
		}
	}
	for _,v:=range chal{
		go output(v);
	}
	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	c1 := generator(1, 5)
	c2 := generator(7, 110)

	s1 := cube(c1)
	s2 := cube(c2)
	res := merger(s1, s2)

	for r:=range res{
			fmt.Println(r)
			
		}
	// var wg sync.WaitGroup
	// wg.Add(100)

	// print:=func (val int)  {
	// 	defer wg.Done()
	// 	fmt.Println(val)
	// }
	// go func ()  {
		
	// 	for r:=range res{
	// 		go print(r);
			
	// 	}
	// }()
	
		
	// 	wg.Wait()
	
	
}