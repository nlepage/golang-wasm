package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg  sync.WaitGroup
	ch1 = make(chan int)
	ch2 = make(chan int)
	end = make(chan struct{})
)

func main() {
	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
		close(out)
	}(ch1)

	go func(in <-chan int, out chan<- int) {
		for i := range in {
			out <- i * i
		}
		close(out)
	}(ch1, ch2)

	go func(in <-chan int) {
		for i := range in {
			fmt.Println(i)
		}

		end <- struct{}{}
	}(ch2)

	fmt.Println("Hello wasm !")

	<-end
}
