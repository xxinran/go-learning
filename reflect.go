package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	SayHello(name string)
	Run()
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Printf("Hi, %v, I am %v", name, hero.Name)
}

func (hero *Hero) Run() {
	fmt.Printf("I am running at speed %v", hero.Speed)
}

func main() {
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Println("type of & is", typeOfHero, "kind is", typeOfHero.Kind())
	// type of Hero is main.Hero kind is struct
	// type是指变量所属的类型
	// kind是指变量的类型所属的类别
	typeOfpHero := reflect.TypeOf(&Hero{})
	fmt.Println("type of pHero is", typeOfpHero, "kind is", typeOfpHero.Kind())
	// type of pHero is *main.Hero kind is ptr

	var person Person
	fmt.Println(reflect.TypeOf(person)) //nil
	person = &Hero{}
	fmt.Println(reflect.TypeOf(person)) //*main.Hero
	// Method
	fmt.Println(reflect.TypeOf(person).NumMethod())
	for i := 0; i < reflect.TypeOf(person).NumMethod(); i++ {
		fmt.Println(reflect.TypeOf(person).Method(i))
	}
	// Value
	name := "Alice"
	fmt.Println(reflect.ValueOf(name).String())
	fmt.Println(reflect.ValueOf(&name))            //0xc000010280
	reflect.ValueOf(&name).Elem().SetString("Bob") //必须是指针
	fmt.Println(reflect.ValueOf(name).String())

	hero := &Hero{
		Name: "Cindy",
	}
	fmt.Println(reflect.ValueOf(&hero)) // 0xc000102060
	fmt.Println(reflect.ValueOf(hero))  //	&{Cindy 0 0}

	// set value
	valueofhero := reflect.ValueOf(hero)
	valueofname := valueofhero.Elem().FieldByName("Name")
	fmt.Println(valueofhero, valueofname)
	if valueofname.CanSet() {
		valueofname.Set(reflect.ValueOf("David"))
	}
	fmt.Println(valueofname)
}
