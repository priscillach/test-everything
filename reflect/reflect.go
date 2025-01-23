package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func (c Cat) SayHello() {
	fmt.Println("Meow!")
}

func main() {
	//cat := Cat{}
	//v := reflect.ValueOf(cat)
	//t := reflect.TypeOf(cat)
	//fmt.Println(t)
	//f := v.MethodByName("SayHello")
	//f.Call([]reflect.Value{})

	e := "aaa"
	v := reflect.ValueOf(&e)
	v = v.Elem()
	v.SetString("123")
	fmt.Println(e) // Output: 123
}
