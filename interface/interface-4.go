package main

import "fmt"

type I interface {
	Get() int
	Set(v int)
}

type S struct {
	age int
}

type SS struct {
	size int
}

func (s S) Get() int {
	return s.age
}

func (s *S) Set(v int) {
	s.age = v
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())

	if t, ok := i.(*S); ok {
		fmt.Println(t, ok)
	}
	switch i.(type) {
	case (*S):
		fmt.Println("type is *S")
	// case (S):
	// 	fmt.Println("type is S") //不允许判断这个类型，因为实现了接收者是*S的方法，所以不可能是S类型
	default:
		fmt.Println("other type")
	}

}

func main() {
	s := new(S)
	f(s)
	ss := new(SS)
	f(ss)
	// 由于没有为SS类型实现Get和Set接口
	// cannot use ss (type *SS) as type I in argument to f:
	// *SS does not implement I (missing Get method)g
}
