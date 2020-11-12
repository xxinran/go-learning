package main

import "fmt"

// func funct() func() {
// 	name := "xinran"
// 	return func() {
// 		fmt.Println(name)
// 	}
// }

// func funct() func() {
// 	name := "xinran"
// 	func printer() {
// 		fmt.Println(name)
// 	}
// 	return printer
// }
//貌似不能先定义再返回函数名，必须直接返回匿名函数和它的定义
func funct() func() {
	name := "xinran"
	return func() {
		fmt.Println(name)
	}
}

func a() func() int { //注意这里的返回func类型是func() int
	var start = 5
	// func b(){
	// 	start = start + 1
	// }
	return func() int {
		start = start + 1
		return start
	}
}

func main() {
	f := funct()
	// fmt.Printf("%T %v", f, f)
	f()
	//这样main函数就可以

	f1 := a()
	fmt.Println(f1())
	fmt.Println(f1()) // 证明main()函数也可以读到a()函数内部变量start
	fmt.Println(f1()) // 这就是闭包的意义
}
