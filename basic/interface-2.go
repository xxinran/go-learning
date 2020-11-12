package main

import (
	"fmt"
)

type Cat interface {
	CatchMouse()
}

type Dog interface {
	Bark()
}

type CatDog struct {
	Name string
}

// 结构体实现接口
func (catdog CatDog) CatchMouse() {
	fmt.Printf("%v is catching a mouse \n", catdog.Name)
}

func (catdog CatDog) Bark() {
	fmt.Printf("%v is barking \n", catdog.Name)
}

func main() {
	var cd CatDog
	cd.Name = "Alice"
	cd.CatchMouse()
	cd.Bark()

	var cat Cat
	fmt.Printf("%T", cat)
	cat = cd
	fmt.Printf("%T", cat)
	cat.CatchMouse()

	var dog Dog
	dog = cd
	dog.Bark()
}
