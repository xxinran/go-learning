package main

import (
	"fmt"

	"rsc.io/quote"
)

// func main() {
// 	fmt.Println("Hello world")
// 	fmt.Println(quote.Hello())
// 	proc("wangxinran", func(str string) {
// 		fmt.Println(str)
// 	})
// }

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Hello())
	proc("wangxinran", printstr)
	fmt.Println(getodd(4, quadra))
}

func printstr(str string) {
	fmt.Println(str)
}
func proc(input string, processor func(str string)) {
	processor(input)
}

func double(x int) int {
	return x * 2
}

func quadra(x int) int {
	return x * 4
}

func getodd(x int, callback func(x int) int) int {
	// fmt.Println(callback(x))
	return 1 + callback(x)
}
