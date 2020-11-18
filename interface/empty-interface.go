package main

import "fmt"

func printAll(vals []interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}

func main() {
	names := []string{"alex", "bob", "cindy"}
	// printAll(names)
	// 报错：
	// cannot use names (type []string) as type []interface {} in argument to printAll
	// go 不会对类型时interface{}的slice进行转换
	// interface{} 会占用两个字长的存储空间，一个是自身的 methods 数据，一个是指向其存储值的指针
	// 也就是 interface 变量存储的值，因而 slice []interface{} 其长度是固定的N*2
	// 但是 []T 的长度是N*sizeof(T)，两种 slice 实际存储值的大小是有区别的

	// 所以需要手动改进型转换
	// 如下：
	fmt.Println(len(names)) //3
	var interfaceslice []interface{} = make([]interface{}, len(names))
	for i, name := range names {
		interfaceslice[i] = name
	}
	fmt.Println(interfaceslice)
	printAll(interfaceslice)

}
