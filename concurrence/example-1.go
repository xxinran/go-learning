package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("I am work in a single goroutine")
}

func main() {
	go test()
	time.Sleep(2 * time.Second)
	go func() {
		fmt.Println("I am work in a single goroutine, anonynmous func")
	}() // 后面这个括号表示是函数调用
	time.Sleep(2 * time.Second)
	go func(str string) {
		fmt.Println("I am work in a single goroutine, anonynmous func", str)
	}("with param") // 后面这个括号表示是函数调用
	time.Sleep(2 * time.Second)
	// 加sleep的原因：
	// 如果不加，大多数情况test中打印操作不会被执行
	// 因为main也是一个goroutine，只要主goroutine结束了，那么整个程序就结束了。
	// 如果 go test()没有在主goroutine之前被调度到，那么就不会被执行
}
