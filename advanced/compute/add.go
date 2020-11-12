package compute

// type AddOperator interface {
// 	Add()
// }

//实现或不实现上面的interface都可以成功运行

type IntParam struct {
	P1 int //大写保证包外也可以被调用
	P2 int
}

func (param *IntParam) Add() int {
	return param.P1 + param.P2
}

// func main() {
// 	param := IntParam{1, 3}
// 	fmt.Println(param.Add())
// }

// 在当前目录执行go install会在GOPATH/pkg/linux_amd64/xinran/advanced下生成一个compute.a中间文件
// 因为当前包里没有main
