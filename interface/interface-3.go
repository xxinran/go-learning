package main

import "fmt"
import "reflect"

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	vtype := reflect.TypeOf(v)
	vtypekind := vtype.Kind()
	fmt.Println(v, vtype, vtypekind)
	// <nil>
	// *main.TestStruct
	// ptr
	return v == nil
}

func main() {
	var s *TestStruct
	fmt.Println(s == nil)    // #=> true
	fmt.Println(NilOrNot(s)) // #=> false

}
