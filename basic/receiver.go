package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { //非指针类别接收器
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { //指针类别接收器
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)          //指针类型接收器，更改的是v本身的值，执行完毕后v被修改。
	fmt.Println(v.Abs()) //不需要修改v的值，可以传一个值拷贝（非指针类型）。
}
