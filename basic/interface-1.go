package main

import "fmt"

type iGreeting interface {
	sayHello()
}

func sayHello(i iGreeting) {
	i.sayHello()
}

type golang struct{}

func (g golang) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type php struct{}

func (p php) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

func main() {
	golang := golang{}
	php := php{}

	golang.sayHello() //普通的结构体函数调用
	sayHello(php)     // 使用interface
}
