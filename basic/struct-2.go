package main

import "fmt"

type Swimming struct {
}

func (swimming *Swimming) swim() {
	fmt.Println("I can swim.\n")
}

type Flying struct {
}

func (flying *Flying) fly() {
	fmt.Println("I can fly.\n")
}

type WildDuck struct {
	Swimming        //只填类型，称为匿名字段/类型嵌套
	f        Flying //Flying类型的变量
	// Flying
}

type DomesticDuck struct {
	Swimming //只填类型
}

func main() {
	wild := WildDuck{}
	domestic := DomesticDuck{}
	wild.f.fly() //指定名字就需要通过名字f来调用fly()
	// wild.fly()   // 如果是L22，只指定类型
	wild.swim()
	// domestic.fly() //domestic.fly undefined
	domestic.swim()
}
