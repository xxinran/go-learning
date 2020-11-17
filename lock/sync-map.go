package main

import "fmt"
import "sync"

var syncMap sync.Map
var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go addNumber(i * 10)
	}
	waitGroup.Wait()
	var size int
	syncMap.Range(func(key, value interface{}) bool {
		size++
		fmt.Println("key-value is ", key, value)
		return true
	})
	fmt.Println("wait finished")
	fmt.Println("size:", size)
}

func addNumber(begin int) {
	for i := begin; i < begin+3; i++ {
		syncMap.Store(i, i)
	}
	waitGroup.Done()
}
