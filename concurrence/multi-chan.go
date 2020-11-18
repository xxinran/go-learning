package main

import (
	"fmt"
	"time"
)

// Can not figure it out.
// Got read val 2 from chl before put val 2 in channel.
// https://stackoverflow.com/questions/64869934/run-goroutine-with-buffered-channel-i-got-an-output-before-input-an-element-int
func send(ch chan int, size int) {
	for i := size; i < size+10; i++ {
		ch <- i
		fmt.Printf("put val %v in channel\n", i)
	}
	// close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// defer close(ch1)
	// defer close(ch2)

	go send(ch1, 0)
	go send(ch2, 20)

	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Printf("read val %v from chl\n", val)
		case val := <-ch2:
			fmt.Printf("read val %v from ch2\n", val)
		case <-time.After(2 * time.Second): //这里是从阻塞的channel中超时返回的小技巧
			fmt.Println("Time out")
			return
		}
	}

}

// func write(ch chan int) {
// 	for i := 0; i < 4; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
// func main() {

// 	// creates capacity of 2
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)

// 	}
// }
