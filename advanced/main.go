package main

import (
	"fmt"
	"std/xinran/advanced/compute" //vscode自动生成的
)

// 需要同时go run main.go hello.go
// 否则 go run main.go 报错 ./main.go:4:2: undefined: sayHello
func main() {
	sayHello()
	param := compute.IntParam{1, 3}
	fmt.Println(param.Add())
}

// 在当前目录下执行go install会生成一个名为advanced的可执行文件在GOPATH/bin目录下
