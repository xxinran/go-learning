package main

import "sync"

var a string
var wg sync.WaitGroup

func hello() {
	go func() {
		a = "hello"
		print("in go")
		wg.Done()
	}()
	print("out of go")
	// wg.Done()
}

func main() {
	wg.Add(1)
	hello()
	wg.Wait()
}
