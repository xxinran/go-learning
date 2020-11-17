package main

import "fmt"

type Sample struct{}

func (s Sample) DoA() {
	fmt.Println("I am DoA()")
}
func (s *Sample) DoB() {
	fmt.Println("I am DoB()")
}

func main() {
	var s Sample
	var sp *Sample

	s.DoA()
	s.DoB()
	// sp.DoA() //会报错，原因？？？
	sp.DoB()
}
