package main

import (
	"fmt"
)

type person struct {
	name string
	id   int
	age  int
}

func main() {
	// 实例化，分配内存空间
	var p1 person
	p2 := new(person)
	p3 := &person{}
	// 初始化，赋值
	p1.name = "a"
	p1.age = 1
	p1.id = 001
	p2.name = "b"
	p2.age = 2
	p2.id = 002
	p3.name = "c"
	p3.age = 3
	p3.id = 003
	p4 := person{
		name: "d",
		age:  4,
		id:   004, //最后一项后面也要加逗号
	}
	fmt.Printf("p1 type %T\n p2 type %T\n p3 type %T\n", p1, p2, p3)
	fmt.Println(p1, *p2, *p3, p4)

}
