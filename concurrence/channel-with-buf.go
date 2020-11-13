package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	time.Sleep(time.Second * 10)

	fmt.Println("consuming", <-ch)

}

func main() {
	ch := make(chan int, 2)
	defer close(ch)
	go consume(ch)
	ch <- 0
	ch <- 1
	fmt.Println("I am free")
	ch <- 2
	fmt.Println("I go to channel after 0 is consumed")
	time.Sleep(time.Second)
	fmt.Println(ch)

}
