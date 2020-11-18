package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go loop()
	//time.Sleep(time.Second)
	for i := 10; i < 20; i++ {
		//go loop()
		fmt.Println(i)
	}
	time.Sleep(time.Second)
}
